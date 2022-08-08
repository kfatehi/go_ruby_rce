require 'ripper'
require 'json'
sexp = Ripper.sexp(ARGF)
raise "could not parse ruby" unless sexp

# scan top-level functions for one called run

top_run_funcs = sexp[1].select do |s|
    s[0] == :def && s[1][1] == "run"
end

if top_run_funcs.size == 0
    raise "no run function detected at top level"
end
# get_run_func_params(top_run_funcs.last)
rf = top_run_funcs.last
param_section = rf[2]
# now we can have paren or not.
if param_section[0] == :paren
    param_list = param_section[1]
else
    param_list = param_section
end
# now it doesnt matter if you have paren
# now lets deal with labels and defaults
def extract_value a
    # puts a.inspect
    case a[0]
    when :var_ref
        extract_value(a[1])
    when :@kw
        a[1]
    when :string_literal
        extract_value(a[1])
    when :string_content
        extract_value(a[1])
    when :@tstring_content
        a[1]
    when :array
        extract_value(a[1])
    else
        # puts "???? #{a}"
    end
end
def crawl_tree(node, result={})
    last = nil
    node.each_with_index do |n, i|
        node_class = n[0].class.to_s
        if node_class == "Array"
            crawl_tree(n, result)
        else
            if i == 0
                if n[0] == :@ident
                    result[n[1]] = {type: :variable}
                elsif n[0] == :@label
                    result[n[1]] = {type: :label}
                end
                last = n[1]
            elsif i == 1
                case result[last][:type]
                when :variable
                    result[last][:value] = extract_value(n)
                when :label
                    result[last][:value] = extract_value(n[1])
                end
            end
            # puts "#{i} #{n}"
        end
    end
    result
end
ordered_params_hash = crawl_tree(param_list[1..-1].compact)

# puts ordered_params_hash

# so as we can see, it's quite complex to try and get default values.
# instead, we will format our final output not with default values
# but with a flag for whether or not some input is required

output = []
ordered_params_hash.each do |k,v|
    output << {name: k, required: !v.has_key?(:value)}
end
puts JSON.generate output
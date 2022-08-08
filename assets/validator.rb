require 'ripper'
require 'json'
sexp = Ripper.sexp(ARGF)
raise "could not parse ruby" unless sexp

result = {


}

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
def crawl_tree(node)
    node.each_with_index do |n, i|
        node_class = n[0].class.to_s
        if node_class == "Array"
            crawl_tree(n)
        else
            puts "#{i} #{n}"
        end
    end
end
crawl_tree(param_list[1..-1].compact)


# pp sexp
# puts JSON.generate result
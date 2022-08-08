package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets15e59ec61440888368f15c5f035ea96a0af5db02 = "require 'ripper'\nrequire 'json'\nsexp = Ripper.sexp(ARGF)\nraise \"could not parse ruby\" unless sexp\n\n# scan top-level functions for one called run\n\ntop_run_funcs = sexp[1].select do |s|\n    s[0] == :def && s[1][1] == \"run\"\nend\n\nif top_run_funcs.size == 0\n    raise \"no run function detected at top level\"\nend\n# get_run_func_params(top_run_funcs.last)\nrf = top_run_funcs.last\nparam_section = rf[2]\n# now we can have paren or not.\nif param_section[0] == :paren\n    param_list = param_section[1]\nelse\n    param_list = param_section\nend\n# now it doesnt matter if you have paren\n# now lets deal with labels and defaults\ndef extract_value a\n    # puts a.inspect\n    case a[0]\n    when :var_ref\n        extract_value(a[1])\n    when :@kw\n        a[1]\n    when :string_literal\n        extract_value(a[1])\n    when :string_content\n        extract_value(a[1])\n    when :@tstring_content\n        a[1]\n    when :array\n        extract_value(a[1])\n    else\n        # puts \"???? #{a}\"\n    end\nend\ndef crawl_tree(node, result={})\n    last = nil\n    node.each_with_index do |n, i|\n        node_class = n[0].class.to_s\n        if node_class == \"Array\"\n            crawl_tree(n, result)\n        else\n            if i == 0\n                if n[0] == :@ident\n                    result[n[1]] = {type: :variable}\n                elsif n[0] == :@label\n                    result[n[1]] = {type: :label}\n                end\n                last = n[1]\n            elsif i == 1\n                case result[last][:type]\n                when :variable\n                    result[last][:value] = extract_value(n)\n                when :label\n                    result[last][:value] = extract_value(n[1])\n                end\n            end\n            # puts \"#{i} #{n}\"\n        end\n    end\n    result\nend\nordered_params_hash = crawl_tree(param_list[1..-1].compact)\n\n# puts ordered_params_hash\n\n# so as we can see, it's quite complex to try and get default values.\n# instead, we will format our final output not with default values\n# but with a flag for whether or not some input is required\n\noutput = []\nordered_params_hash.each do |k,v|\n    output << {name: k, required: !v.has_key?(:value)}\nend\nputs JSON.generate output"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"assets"}, "/assets": []string{"validator.rb"}}, map[string]*assets.File{
	"/assets/validator.rb": &assets.File{
		Path:     "/assets/validator.rb",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1659922112, 1659922112004665729),
		Data:     []byte(_Assets15e59ec61440888368f15c5f035ea96a0af5db02),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1659921764, 1659921764306842331),
		Data:     nil,
	}, "/assets": &assets.File{
		Path:     "/assets",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1659921334, 1659921334647550615),
		Data:     nil,
	}}, "")

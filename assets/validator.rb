require 'ripper'
require 'json'
sexp = Ripper.sexp(ARGF)
puts JSON.generate sexp
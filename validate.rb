require 'ripper'
require 'json'

f = File.read(ENV['SCRIPT_FILE'])
puts JSON.generate Ripper.sexp(f)
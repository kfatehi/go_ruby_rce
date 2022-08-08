require 'foo'
require 'bar'
require 'baz'

def run
    puts "oh i will overwrite this later!"
end

run

def run foo, some_arr=[1,2,"chaos"], bar="hi", baz={ping: pong}, dry_run: true, other_thing: nil, hi: "bye", &block
    puts "some stuff"
end

run
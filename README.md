# go-ruby-rce

A remote (ruby) code execution agent (webserver) written in Go.

INCOMPLETE/WIP/ABANDONED BRAINSTORMING PROJECT

## Disclaimer
**This dangerous project is intended to run on a private network behind a firewall!
Only authorized users should be able to cause interactions with this program!
I am not responsible for any damages should you install this.**

## Overview & Reason for Existence

This program provides a simple API for validation and execution of ruby scripts.
It is designed to facilitate script execution for Ruby on Rails environments where
ruby scripts are often built in response to support requests. Such scripts are not
written to the standard of quality as required for the application itself, and upon
successful execution, are typically lost in the trashbin of the author's harddrive,
unorganized. This leads to even more developer time being wasted writing the same
scripts. This is not even considering the time wasted reviewing the code that will
be executing in production -- a crucial chore with serious consequences.

If developers organized their scripts collectively, and the quality (success rate)
were to be measured over time, the pain will significantly decrease. Confidence about
particular tasks will increase. Time wasted writing and reviewing scripts will decrease.

Furthermore, when input parameters of these scripts are parsed and used to generate a
form input, significant tedium can be pushed to the support staff, empowering both the
support and the development personnel. Customers, developers, and support staff are happy!

Win win win!

## API Interface

The main web application should use it to do the following (via protected routes):

### POST /ruby/validate

Analyze and validate a script, pulling out a list of parameters.

You should POST a ruby script with a run function at the top level using multipart forms, e.g.:

```
cat <<EOF > my_ruby_script.rb
def run foo, a=nil, some_arr=[1,2,"chaos"], bar="hi", baz={ping: pong}, dry_run: true, other_thing: nil, hi: "bye", &block
    puts "the actual work...
end
EOF
curl localhost:8080/ruby/validate -F file=@my_ruby_script.rb
```

Internally, ruby's `Ripper` will check the script for validity (has a top-level run function) and return a JSON array with the arguments (and their optionality) for use in form generation.

### Success Criteria

Success looks like `200 OK` `application/json` JSON with a top-level array object. In our above example, you would see:

```
[{"name":"foo","required":true},{"name":"a","required":false},{"name":"some_arr","required":false},{"name":"bar","required":false},{"name":"baz","required":false},{"name":"dry_run:","required":false},{"name":"other_thing:","required":false},{"name":"hi:","required":false}]
```

### Failure Criteria

Failures looks like `>400` `application/json` JSON like so:

```
{"error":"..."}
```

### POST /ruby/execute

execute a script with provided arguments and environment variables.
effectively a remote spawn function
a unique id for the job will be returned to you.
you keep track of it

### GET /ruby/jobs/:id

obtain exit statuses and logs for the given job

## Expected Outer Application Responsibilities

It remains the main application's responsibility to:

1. Organize scripts
2. Utilize validation API to guard script creation to ensure validity
3. Utilize validation API parameter analysis response during script creation to generate nice forms
4. Organize targets (machine(s) which run this program)
5. Keep track of job IDs provided by the execution API
5. Utilize the Jobs API to handle outcomes

All data is stored on the filesystem in the configured locations. The sysadmin is responsible for clearing this data over time.

## Configuration

Configurations are provided via environment variables

1. RUBY - Path to ruby executable (for using Ripper), defaults to whatever your PATH provides.
1. LISTEN_ADDR - Address to bind to, defaults to localhost
1. LISTEN_PORT - Port to bind to, defaults to 8080
2. TBD... Path to for file storage 

---

Future consideration:

- https://github.com/distribworks/dkron Distributed, fault tolerant job scheduling system

actually, using dkron for the execution system might be pay off nicely in not having to write the more choreish parts of the next two APIs (execute and status).

https://github.com/go-co-op/gocron

deps:
	go get
	go get github.com/jessevdk/go-assets-builder

assets:
	go run github.com/jessevdk/go-assets-builder assets -o assets.go

run:
	go run .

build:
	go build

test:
	go test

ruby:
	cat test_support/testscript.rb | ruby assets/validator.rb

watchruby:
	fswatch -o --event Updated assets/*.rb test_support/*.rb | xargs -n1 -I{} bash -c "clear; just ruby"

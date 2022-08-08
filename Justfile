assets:
	go run github.com/jessevdk/go-assets-builder assets -o assets.go

start: assets
	go run .

test: assets
	go test

build: assets
	go build

watchruby:
	fswatch -o --event Updated assets/*.rb test_support/*.rb | xargs -n1 -I{} bash -c "clear; cat test_support/testscript.rb | ruby assets/validator.rb"

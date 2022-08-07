assets:
	go run github.com/jessevdk/go-assets-builder assets -o assets.go

start: assets
	go run .

test: assets
	go test

build: assets
	go build
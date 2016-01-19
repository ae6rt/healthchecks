all: build

lint:
	godep go vet ./...

test: lint
	godep go test -v ./...

build: test
	godep go build

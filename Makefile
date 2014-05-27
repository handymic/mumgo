export GOPATH = $(shell pwd)

all: fmt test

fmt:
	for f in `ls *.go`; do \
		gofmt -tabwidth=2 -w $$f; done

test:
	go test

export GOPATH = $(shell pwd)

all: fmt

fmt:
	for f in `ls *.go`; do \
		gofmt -tabwidth=2 -w $$f; done


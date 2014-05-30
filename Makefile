export GOPATH = $(shell pwd)
export TMPDIR = $(shell pwd)/tmp/mumgo
export TEST_CRT = $(GOPATH)/tmp/certs/mumgo.crt
export TEST_KEY = $(subst .crt,.key,$(TEST_CRT))


all: fmt test

fmt:
	for f in `ls *.go`; do \
		gofmt -tabwidth=2 -w $$f; done

test: $(TMPDIR) $(TEST_CRT) setup
	go test

setup:
	go get -d .

$(TEST_CRT):
	# Generating test cert required for tests ...
	mkdir -p $(shell dirname $(TEST_CRT)) && \
		openssl req -x509 -newkey rsa:2048 -nodes -days 1001 \
			-keyout $(TEST_KEY) -out $@

$(TMPDIR):
	mkdir -p $@

export GOPATH = $(shell pwd)
export TEST_CRT = $(GOPATH)/tmp/certs/mumgo.crt
export TEST_KEY = $(subst .crt,.key,$(TEST_CRT))


all: fmt test

fmt:
	for f in `ls *.go`; do \
		gofmt -tabwidth=2 -w $$f; done

test: $(TEST_CRT)
	go test

$(TEST_CRT):
	# Generating test cert required for tests ...
	mkdir -p $(shell dirname $(TEST_CRT)) && \
		openssl req -x509 -newkey rsa:2048 -nodes -days 1001 \
			-keyout $(TEST_KEY) -out $@


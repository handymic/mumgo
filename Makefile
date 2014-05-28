export GOPATH = $(shell pwd)
TEST_CERT = $(GOPATH)/tmp/certs/mumgo.crt


all: fmt test

fmt:
	for f in `ls *.go`; do \
		gofmt -tabwidth=2 -w $$f; done

test: $(TEST_CERT)
	go test

$(TEST_CERT):
	# Generating test cert required for tests ...
	mkdir -p $(shell dirname $(TEST_CERT)) && \
		openssl req -x509 -newkey rsa:2048 -nodes -days 1001 \
			-keyout $(subst .crt,.key,$@) -out $@


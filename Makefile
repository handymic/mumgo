export ROOT = $(shell pwd)
export GOPATH = $(ROOT)
export TMPDIR = $(ROOT)/tmp/mumgo
export TEST_CRT = $(ROOT)/tmp/certs/mumgo.crt
export TEST_KEY = $(subst .crt,.key,$(TEST_CRT))

MURMUR_CFG = $(ROOT)/tmp/murmur/murmur/murmur.ini


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


run/deps: 
	mkdir -p $(shell dirname $(MURMUR_CFG)) && \
		sed "s|@@ROOT@@|$(ROOT)|g" \
			murmur/$(shell basename $(MURMUR_CFG)) > $(MURMUR_CFG) && \
				murmurd -fg -ini $(MURMUR_CFG)


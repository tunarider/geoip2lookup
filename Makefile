SHELL := /bin/bash
DEST := /usr/local/bin

.EXPORT_ALL_VARIABLES:

GOBIN = $(shell pwd)/bin

geoip2lookup: dep
	cd ./cmd/geoip2lookup && go install

dep:
	dep ensure

install:
	cp $(GOBIN)/geoip2lookup $(DEST)

clean:
	rm -rf $(GOBIN)/bin

uninstall:
	rm $(DEST)/geoip2lookup

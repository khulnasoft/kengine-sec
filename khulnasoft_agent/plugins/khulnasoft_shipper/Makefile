.PHONY: khulnasoft_shipper

VERSION?=v`git rev-parse --short HEAD`

all: khulnasoft_shipper

local: khulnasoft_shipper

vendor: go.mod $(shell find ../../../khulnasoft_utils -path ../../../khulnasoft_utils/vendor -prune -o -name '*.go')
	go mod tidy -v
	go mod vendor

khulnasoft_shipper: vendor $(shell find . -path ./vendor -prune -o -name '*.go')
	CGO_ENABLED=0 go build -buildvcs=false -ldflags="-s -w -X main.Version=${VERSION} -X main.Commit=${VERSION} -X main.BuildTime=$(shell date +"%s%d%m%y") -extldflags=-static"

clean:
	-rm khulnasoft_shipper
	-rm -rf ./vendor

.PHONY: all clean local

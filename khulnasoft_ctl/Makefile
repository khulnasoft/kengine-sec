all: khulnasoftctl

vendor: go.mod $(shell find ../khulnasoft_utils -path ../khulnasoft_utils/vendor -prune -o -name '*.go')
	go mod tidy -v
	go mod vendor

khulnasoftctl: vendor $(shell find . -path ./vendor -prune -o -name '*.go')
	CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=${VERSION} -extldflags=-static" -buildvcs=false -o khulnasoftctl

clean:
	-rm khulnasoftctl

.PHONY: all clean

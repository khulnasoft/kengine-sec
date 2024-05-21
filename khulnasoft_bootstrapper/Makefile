VERSION?=v`git describe --tags`

all: khulnasoft_bootstrapper

local: khulnasoft_bootstrapper

prepare:
	docker run --rm -i -e VERSION=${VERSION} -v $(ROOT_MAKEFILE_DIR):/src:rw -v /tmp/go:/go:rw $(IMAGE_REPOSITORY)/khulnasoft_builder_ce:$(KE_IMG_TAG) bash -c 'cd /src/khulnasoft_bootstrapper && make khulnasoft_bootstrapper'

vendor: go.mod $(shell find ../khulnasoft_utils -path ../khulnasoft_utils/vendor -prune -o -name '*.go')
	go mod tidy -v
	go mod vendor

khulnasoft_bootstrapper: vendor $(shell find . -path ./vendor -prune -o -name '*.go') $(shell find ./assets/ -name '*.ini')
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -X main.Version=${VERSION} -extldflags=-static" -buildvcs=false
	-upx ./khulnasoft_bootstrapper

clean:
	-rm khulnasoft_bootstrapper
	-rm -rf ./vendor

.PHONY: clean image local all

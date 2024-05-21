VERSION?=$(shell git describe --tags)
GIT_COMMIT=$(shell git rev-parse HEAD)
BUILD_TIME=$(shell date -u +%FT%TZ)

all: khulnasoft_server

local: khulnasoft_server

image:
	docker run --rm -i -e VERSION=${VERSION} -e GIT_COMMIT=${GIT_COMMIT} -e BUILD_TIME=${BUILD_TIME} -v $(ROOT_MAKEFILE_DIR):/src:rw -v /tmp/go:/go:rw $(IMAGE_REPOSITORY)/khulnasoft_glibc_builder_ce:$(KE_IMG_TAG) bash -c 'cd /src/khulnasoft_server && make khulnasoft_server'
	docker build -f ./Dockerfile -t $(IMAGE_REPOSITORY)/khulnasoft_server_ce:$(KE_IMG_TAG) ..

vendor: go.mod $(shell find ../khulnasoft_utils -path ../khulnasoft_utils/vendor -prune -o -name '*.go')
	go mod tidy -v
	go mod vendor

khulnasoft_server: vendor $(shell find . -path ./vendor -prune -o -name '*.go')
	go build -buildvcs=false -ldflags="-s -w -X github.com/khulnasoft/kengine/khulnasoft_server/pkg/constants.Version=${VERSION} -X github.com/khulnasoft/kengine/khulnasoft_server/pkg/constants.Commit=${GIT_COMMIT} -X github.com/khulnasoft/kengine/khulnasoft_server/pkg/constants.BuildTime=${BUILD_TIME}"

clean:
	-rm khulnasoft_server
	-rm -rf ./vendor

.PHONY: all clean image local

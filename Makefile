MAKEFLAGS   += --warn-undefined-variables
GOOS        ?= $(shell go env GOOS)
GOARCH      ?= $(shell go env GOARCH)
CGO_ENABLED ?= $(shell go env CGO_ENABLED)

cmd/tool/encrypt: cmd/tool/main.go
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED} go build -ldflags="-s -w" -trimpath -o $@ $^

docs/encrypt.wasm: cmd/wasm/main.go
	GOOS=js GOARCH=wasm CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o $@ $^

test:
	@go clean -testcache
	@go test -race ./...

lint:
	@go vet ./...

clean:
	@rm -rf cmd/tool/encrypt

.PHONY: cmd/tool/encrypt docs/encrypt.wasm test lint clean

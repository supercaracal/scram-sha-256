MAKEFLAGS   += --warn-undefined-variables
GOOS        ?= $(shell go env GOOS)
GOARCH      ?= $(shell go env GOARCH)
CGO_ENABLED ?= $(shell go env CGO_ENABLED)

define go-build
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED}\
	go build -ldflags="-s -w" -trimpath -o $@ $^
endef

build: term wasm

term: cmd/term/encrypt

wasm: docs/wasm_exec.js docs/encrypt.wasm cmd/debug/server

test:
	@go clean -testcache
	@go test -race ./...

lint:
	@go vet ./...

clean:
	@rm -f cmd/term/encrypt cmd/debug/server

cmd/term/encrypt: cmd/term/main.go
	$(call go-build)

cmd/debug/server: cmd/debug/main.go
	$(call go-build)

docs/encrypt.wasm: GOOS        := js
docs/encrypt.wasm: GOARCH      := wasm
docs/encrypt.wasm: CGO_ENABLED := 0
docs/encrypt.wasm: cmd/wasm/main.go
	$(call go-build)

docs/wasm_exec.js: $(shell go env GOROOT)/misc/wasm/wasm_exec.js
	@cp $^ $@

.PHONY: build term wasm test lint clean \
	cmd/term/encrypt docs/encrypt.wasm

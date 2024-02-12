MAKEFLAGS   += --warn-undefined-variables
SHELL       ?= /bin/bash -euo pipefail
GOOS        ?= $(shell go env GOOS)
GOARCH      ?= $(shell go env GOARCH)
CGO_ENABLED ?= $(shell go env CGO_ENABLED)

define go-build
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED}\
	go build -ldflags="-s -w" -trimpath -o $@ $^
endef

all: build test lint

build: term wasm

term: cmd/term/encrypt

wasm: docs/wasm_exec.js docs/encrypt.wasm cmd/debug/server

test:
	@go clean -testcache
	@go test -race ./...

lint:
	@go vet ./...

bench:
	@go test -bench=. -benchmem -run=NONE ./...

prof: PKG ?= pgpasswd
prof: TYPE ?= mem
prof:
	@if [ -z "${PKG}" ]; then echo 'empty variable: PKG'; exit 1; fi
	@if [ -z "${TYPE}" ]; then echo 'empty variable: TYPE'; exit 1; fi
	@if [ ! -d "./pkg/${PKG}" ]; then echo 'package not found: ${PKG}'; exit 1; fi
	@go test -bench=. -run=NONE -${TYPE}profile=${TYPE}.out ./pkg/${PKG}
	@go tool pprof -text -nodecount=10 ${PKG}.test ${TYPE}.out

clean:
	@rm -f cmd/term/encrypt cmd/debug/server *.test *.out

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

.PHONY: build term wasm test lint bench prof clean \
	cmd/term/encrypt docs/encrypt.wasm

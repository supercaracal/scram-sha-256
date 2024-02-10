MAKEFLAGS   += --warn-undefined-variables
GOOS        ?= $(shell go env GOOS)
GOARCH      ?= $(shell go env GOARCH)
CGO_ENABLED ?= $(shell go env CGO_ENABLED)

cmd/tool/encrypt: cmd/tool/main.go
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED} go build -ldflags="-s -w" -trimpath -o $@ $^

.PHONY: cmd/tool/encrypt

include .env

ifdef GOROOT
GO := $(GOROOT)/bin/go
else
GO := go
endif

.PHONY: server
server:
	$(GO) build -o ./bin/server ./cmd/server

.PHONY: merge-profiles
merge-profiles:
	$(GO) tool pprof -proto pprof/*/cpu.pprof > merged.pprof

.PHONY: server-pgo
server-pgo:
	$(GO) build -o ./bin/server -pgo=./cpu.pprof ./cmd/server

.PHONY: server-cover
server-cover:
	$(GO) build -cover -o ./bin/server ./cmd/server

.PHONY: cover-convert
cover-convert:
	$(GO) tool covdata textfmt -i coverage -o coverage.out

.PHONY: slices
slices:
	$(GO) build -o ./bin/slices ./cmd/slices

.PHONY: errors
errors:
	$(GO) build -o ./bin/errors ./cmd/errors


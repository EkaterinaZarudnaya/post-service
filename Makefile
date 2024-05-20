GO = go
GOLANGCILINT = golangci-lint

all: lint test

.PHONY: run
run:
	$(GO) run cmd/main.go

.PHONY: build
build:
	$(GO) build -o bin/post-service cmd/*.go

.PHONY: lint
lint:
	$(GOLANGCILINT) run --out-format=github-actions -- ./...

PHONY: test
test:
	$(GO) test -v ./...

.PHONY: test-coverage
test-coverage:
	$(GO) test -v -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

.PHONY: serve
serve:
	$(GO) run cmd/*.go

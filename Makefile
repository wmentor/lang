.PHONY: all lint test

all: test lint

lint:
	@if [ -x "$$(command -v golangci-lint)" ]; then echo "run golangci-lint..." ; golangci-lint run ; else echo "golangci-lint not found" ; fi

test:
	@echo "run tests..."
	@go clean -testcache
	@go test ./... -cover -race

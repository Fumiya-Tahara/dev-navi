.PHONY: build test lint
GOLANGCI_LINT_VERSION:=v1.57.2-alpine

build:
	@go build -v ./...

test:
	@go test -v ./...

lint:
#	@docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:${GOLANGCI_IMAGE_TAG}
	@docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.57.2-alpine
	@golangci-lint run --config .golangci.yml

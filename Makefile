all: run

deps:dependencies
	@go mod tidy

lint:
	@go vet ./...

run:
	@go run cmd/main.go

test:
	@go test -v ./...

clean-cache:
	@go clean -cache
	@go clean -testcache
	@go clean -modcache

.PHONY: run
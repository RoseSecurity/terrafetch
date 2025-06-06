BINARY_NAME=terrafetch
VERSION=local
GO=go

default: help

help: ## List Makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

all: build

fmt: ## Format Go files
	gofumpt -w .

build: ## Build Terrafetch
	env $(if $(GOOS),GOOS=$(GOOS)) $(if $(GOARCH),GOARCH=$(GOARCH)) $(GO) build -o build/$(BINARY_NAME) main.go

deps: ## Download dependencies
	go mod download

get: ## Install dependencies
	go get

clean: ## Clean up build artifacts
	$(GO) clean
	rm ./build/$(BINARY_NAME)

generate: ## Generate tooling licensing
	cd tools; go generate ./...

testacc: ## Run acceptance tests
	go test ./...

run: build ## Run Terrafetch
	./build/$(BINARY_NAME)

docs: build ## Generate documentation
	./build/$(BINARY_NAME) docs

version: build ## View binary version
	chmod +x ./build/$(BINARY_NAME)
	./build/$(BINARY_NAME) version

.PHONY: all build install clean run fmt help

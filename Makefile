.DEFAULT_GOAL := help
UNAME_S := $(shell uname -s 2>/dev/null || echo "unknown")
UNAME_S_LOWERCASE := $(shell echo $(UNAME_S) | tr A-Z a-z)

.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: deps
deps: ## Install dependencies
ifeq ($(UNAME_S),Linux)
	sudo apt-get update && sudo apt-get install -y golang sqlite3 golangci-lint redis-server
endif
ifeq ($(UNAME_S),Darwin)
	brew install go sqlite golangci-lint redis
endif
ifeq ($(OS),Windows_NT)
	choco install golang sqlite golangci-lint redis
endif

.PHONY: lint
lint: ## Lint the Go code using golangci-lint
	golangci-lint run --skip-dirs externals,examples

.PHONY: test
test: ## Run tests
	go test -v -cover ./...

.PHONY: generate
generate: ## Run tests
	go generate ./...

.PHONY: benchmark
benchmark: ## Run benchmarks
	go test -v -bench . -benchmem ./... > benchmark.txt

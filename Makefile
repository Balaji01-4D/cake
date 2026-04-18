SHELL := /bin/sh

BINARY := $(notdir $(CURDIR))
BIN_DIR := ./bin

.DEFAULT_GOAL := help

.PHONY: help all build run test clean tidy deps-update deps fmt check precommit

help: ## Show available targets
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z0-9_.-]+:.*##/ {printf "  %-14s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

all: build ## Build the project binary

build: ## Build binary into ./bin
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY) .

run: ## Run the project locally
	go run .

test: ## Run all tests
	go test ./...

clean: ## Remove build artifacts (./bin)
	rm -rf $(BIN_DIR)

tidy: ## Tidy go modules
	go mod tidy

deps-update: ## Update all dependencies, then tidy go.mod/go.sum
	go get -u -t ./...
	go mod tidy

deps: deps-update ## Alias for deps-update

fmt: ## Run gofumpt only on changed or untracked Go files
	@files="$$( \
		{ \
			git diff --name-only -- '*.go'; \
			git diff --cached --name-only -- '*.go'; \
			git ls-files --others --exclude-standard -- '*.go'; \
		} | sort -u \
	)"; \
	if [ -z "$$files" ]; then \
		echo "No changed or untracked Go files to format."; \
		exit 0; \
	fi; \
	echo "$$files" | xargs -r gofumpt -w

check: fmt test ## Format changed files and run tests

precommit: check build ## Common pre-commit checks

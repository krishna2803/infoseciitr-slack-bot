.PHONY: build format help lint 

.DEFAULT_GOAL := help

GO := go
GOPATH := $(shell go env GOPATH)
GOPATH_BIN := $(GOPATH)/bin
AIR := $(GOPATH_BIN)/air
GOLANGCI_LINT := $(GOPATH_BIN)/golangci-lint
GO_PACKAGES := $(shell go list ./... | grep -v vendor)
BUILD_OUTPUT := ./target/slack-bot
BUILD_INPUT := cmd/main.go
UNAME := $(shell uname)

all: lint build

help:
	@echo "SlackBot Makefile"
	@echo "build   - Build slack-bot"
	@echo "dev     - Run development environment"
	@echo "format  - Format code using golangci-lint"
	@echo "help    - Prints help message"
	@echo "lint    - Lint code using golangci-lint"

build:
	@echo "Building..."
	@test -d target || mkdir target
	@$(GO) build -o $(BUILD_OUTPUT) $(BUILD_INPUT)
	@echo "Built as $(BUILD_OUTPUT)"

run: build
	@echo "Running..."
	@$(BUILD_OUTPUT)
	
format:
	@echo "Formatting..."
	@$(GO) fmt $(GO_PACKAGES)
	@$(GOLANGCI_LINT) run --fix --issues-exit-code 0 > /dev/null 2>&1
	@echo "Code formatted"

lint:
	@echo "Linting..."
	@$(GO) vet $(GO_PACKAGES)
	@$(GOLANGCI_LINT) run
	@echo "No errors found"

vendor:
	@echo "Tidy up go.mod..."
	@$(GO) mod tidy
	@echo "Vendoring..."
	@$(GO) mod vendor
	@echo "Done!"

install-golangcilint:
	@echo "Installing golangci-lint..."
	@curl -sSfL \
	 	https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
	 	sh -s -- -b $(GOPATH_BIN) v1.43.0
	@echo "Installed successfully"

dev:
	@echo "Starting development server..."
	@$(AIR)


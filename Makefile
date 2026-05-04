.PHONY: build run test test-verbose test-coverage lint fmt clean deps help

# Application name
APP_NAME := server
BINARY_NAME := $(APP_NAME)
BUILD_DIR := ./bin

# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOMOD := $(GOCMD) mod
GOFMT := $(GOCMD) fmt
GOVET := $(GOCMD) vet

# Build flags
LDFLAGS := -s -w
BUILD_FLAGS := -ldflags "$(LDFLAGS)"

# Colors for output
GREEN = \033[0;32m
YELLOW = \033[0;33m
NC = \033[0m # No Color

help: ## Show this help message
	@echo "$(GREEN)Usage: make [target]$(NC)"
	@echo ''
	@echo "$(GREEN)Available targets:$(NC)"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

deps: ## Download dependencies
	@echo "$(GREEN)Downloading dependencies...$(NC)"
	@$(GOMOD) download
	@$(GOMOD) tidy
	@echo "$(GREEN)Dependencies installed$(NC)"

build: ## Build the application
	@echo "$(GREEN)Building $(BINARY_NAME)...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@$(GOBUILD) $(BUILD_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/server
	@echo "$(GREEN)Build complete: $(BUILD_DIR)/$(BINARY_NAME)$(NC)"

run: ## Run the application
	$(GOCMD) run ./cmd/server

test: ## Run tests
	@echo "$(GREEN)Running tests...$(NC)"
	@$(GOTEST) -v ./...

test-verbose: ## Run tests with verbose output
	@echo "$(GREEN)Running tests with verbose output...$(NC)"
	@$(GOTEST) -v -race ./...

test-coverage: ## Run tests with coverage report
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	@$(GOTEST) -v -coverprofile=coverage.out ./...
	@$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

lint: ## Run linters (requires golangci-lint)
	@echo "$(GREEN)Running linter...$(NC)"
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "$(YELLOW)golangci-lint not found. Running go vet instead...$(NC)"; \
		$(GOVET) ./...; \
	fi

fmt: ## Format code
	@echo "$(GREEN)Formatting code...$(NC)"
	@$(GOFMT) ./...
	@echo "$(GREEN)Code formatted$(NC)"

vet: ## Run go vet
	@echo "$(GREEN)Running go vet...$(NC)"
	@$(GOVET) ./...

clean: ## Clean build artifacts
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	@rm -rvf $(BUILD_DIR)
	@rm -vf coverage.out coverage.html server
	@echo "$(GREEN)Clean complete$(NC)"

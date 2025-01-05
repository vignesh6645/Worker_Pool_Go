# Variables
APP_NAME := worker-pool-api
SRC := cmd/main.go
BUILD_DIR := bin

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC)
	@echo "Build complete. Executable: $(BUILD_DIR)/$(APP_NAME)"

# Run the application
.PHONY: run
run: build
	@echo "Running the application..."
	@$(BUILD_DIR)/$(APP_NAME)

# Clean the build directory
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."

# Tidy dependencies
.PHONY: tidy
tidy:
	@echo "Tidying up Go dependencies..."
	@go mod tidy
	@echo "Dependencies are tidy."

# Format code
.PHONY: format
format:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Code formatted."

# Lint code
.PHONY: lint
lint:
	@echo "Linting code..."
	@go vet ./...
	@echo "Linting complete."

# Test the application
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./... -v
	@echo "Tests complete."

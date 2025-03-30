# Define the output directory for binaries
BIN_DIR := bin
GO := go

# Define commands to build
TARGETS := json k8er tray

# Get the OS and ARCH for cross-compilation if needed
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m)

# Default target: Build all binaries
all: $(TARGETS)

# Build each binary separately
$(TARGETS):
	@mkdir -p $(BIN_DIR)
	@echo "Building $@..."
	@$(GO) build -o $(BIN_DIR)/$@ ./cmd/$@/

# Clean built binaries
clean:
	@echo "Cleaning binaries..."
	@rm -rf $(BIN_DIR)

# Run tray (for convenience)
run:
	@$(BIN_DIR)/tray

.PHONY: all clean install run $(TARGETS)

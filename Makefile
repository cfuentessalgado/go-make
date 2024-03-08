.PHONY: build install

GO=go
BINARY_NAME=go-make
BUILD_PATH=./build
BINARY_PATH=$(BUILD_PATH)/$(BINARY_NAME)
INSTALL_PATH=/usr/local/bin

build:
	$(GO) build -o $(BINARY_PATH) ./cmd/go-make

install:
	@if [ -f "$(INSTALL_PATH)/$(BINARY_NAME)" ]; then \
		echo "Existing binary found at $(INSTALL_PATH)/$(BINARY_NAME). Backing up..."; \
		mv $(INSTALL_PATH)/$(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME).bak; \
	fi
	@echo "Copying $(BINARY_NAME) to $(INSTALL_PATH)"
	@cp $(BINARY_PATH) $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "Installation complete."

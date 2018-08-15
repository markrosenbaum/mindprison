GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean

BUILD_ROOT=build
BIN_NAME=mindprison
BUILD_BIN=$(BUILD_ROOT)/$(BIN_NAME)

.PHONY: $(BUILD_BIN)
$(BUILD_BIN):
	$(GOBUILD) -o $(BUILD_BIN) -v ./...

.PHONY: run
run: $(BUILD_BIN)
	$(BUILD_BIN)

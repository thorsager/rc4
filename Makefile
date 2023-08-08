BUILD_DIR := build
BINARY := rc4

.PHONY: build
build:
	go build -o $(BUILD_DIR)/$(BINARY) main.go

clean:
	rm -rf $(BUILD_DIR)

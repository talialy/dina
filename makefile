BIN=bin
DIST_LINUX=$(BIN)/linux
GO_BUILD=go build
BINARY_NAME=momo
GOFMT=gofmt
GOFILES=$(shell find . -name '*.go')


.PHONY: build-linux
build-linux:
	@echo "exporting for linux"
	@mkdir -p $(DIST_LINUX)
	@GOOS=linux GOARCH=amd64 $(GO_BUILD) -o ./$(DIST_LINUX)/$(BINARY_NAME)

.PHONY: test
test: format build-linux
	@./run_tests.sh

.PHONY: format
format:
	@echo "Formatting Go files..."
	$(GOFMT) -w $(GOFILES)
	@echo "Done."

.PHONY: clean
clean:
	@echo "Cleaning up..."
	go clean
	@echo "Cleaned."

.PHONY: all
all: format build-linux

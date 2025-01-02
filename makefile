BIN=bin
DIST=$(BIN)
GO=go
GO_BUILD=$(GO) build
BINARY_NAME=dina
GOFMT=gofmt
GOFILES=$(shell find . -name '*.go')


.PHONY: build
build: format
	@echo "exporting binary"
	@mkdir -p $(DIST)
	@GOOS=linux GOARCH=amd64 $(GO_BUILD) -o ./$(DIST)/$(BINARY_NAME)

.PHONY: test
test: format build
	podman build -t dina .
	podman run -it -e=./.env dina

.PHONY: format
format:
	@echo "Formatting Go files..."
	$(GO) mod tidy
	$(GOFMT) -w $(GOFILES)
	@echo "Done."

.PHONY: clean
clean:
	@echo "Cleaning up..."
	$(GO) clean
	@echo "Cleaned."

.PHONY: all
all: format build

BIN=bin
DIST=$(BIN)
GO_BUILD=go build
BINARY_NAME=dina
GOFMT=gofmt
GOFILES=$(shell find . -name '*.go')


.PHONY: build
build:
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
	$(GOFMT) -w $(GOFILES)
	@echo "Done."

.PHONY: clean
clean:
	@echo "Cleaning up..."
	go clean
	@echo "Cleaned."

.PHONY: all
all: format build

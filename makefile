BIN=bin
DIST_LINUX=$(BIN)/linux
GO_BUILD=go build
BINARY_NAME=momo

build-linux:
	mkdir -p $(DIST_LINUX)
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o ./$(DIST_LINUX)/$(BINARY_NAME)

test:
	build-linux
	./run_tests.sh

PLUGIN_NAME=hc-cache-plugin
BINARY_NAME=hc-cache-plugin
.PHONY: build clean test tidy fmt deps
build:
	go build -o $(BINARY_NAME) .
clean:
	rm -f $(BINARY_NAME)
	go clean -cache
test:
	go test -v ./...
tidy:
	go mod tidy
fmt:
	go fmt ./...
deps:
	go mod download

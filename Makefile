VERSION ?= $(shell git describe --always --tags)
BIN = tbt
BUILD_CMD = go build -o build/$(BIN)-$(VERSION)-$${GOOS}-$${GOARCH} &

default:
	$(MAKE) bootstrap
	$(MAKE) build

test:
	go vet ./...
	golint -set_exit_status $(shell go list ./... | grep -v vendor)
	go test -covermode=atomic -race -v ./...
bootstrap:
	glide install
build:
	go build -o $(BIN)
clean:
	rm -rf build vendor
	rm -f release image bootstrap $(BIN)
release: bootstrap
	@echo "Running build command..."
	bash -c '\
		export GOOS=linux;   export GOARCH=amd64; $(BUILD_CMD) \
		wait \
	'
	touch release

.PHONY: test build clean image-push


NAME := doman
TARGET := $(shell pwd)
BINARY_PATH := $(TARGET)/$(NAME)
GOPATH := $(shell go env GOPATH)

.PHONY: build clean install uninstall

build:
	@go build -o $(BINARY_PATH) ./cmd/doman/main.go

clean:
	@rm -f $(BINARY_PATH)

install:
	@go install ./cmd/doman/

uninstall:
	@rm -f $(GOPATH)/bin/$(NAME)

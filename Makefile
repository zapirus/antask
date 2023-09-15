.PHONY: build
build:
	go build -v ./cmd/task
.DEFAULT_GOAL := build
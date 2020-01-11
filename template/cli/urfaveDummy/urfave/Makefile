.DEFAULT_GOAL: build
GOOS ?= $(shell go env GOOS)
BIN_PATH := bin
build: local
local:
	rm -f $(PWD)/$(BIN_PATH)/main
	GOOS=$(GOOS) go build -o $(PWD)/$(BIN_PATH)/main $(PWD)/main.go
run:
	go run $(PWD)/$(BIN_PATH)/main
docker-build:
	GOOS=$(GOOS) go build -o $(PWD)/$(BIN_PATH)/main $(PWD)/main.go

	docker build -f $(PWD)/Dockerfile -t image:tag $(PWD)

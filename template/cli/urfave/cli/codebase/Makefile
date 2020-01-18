.DEFAULT_GOAL: build
GOOS ?= $(shell go env GOOS)
BIN_PATH := bin
build: local
local:
	rm -f $(PWD)/$(BIN_PATH)/{{ .appname}}
	GOOS=$(GOOS) go build -o $(PWD)/$(BIN_PATH)/{{ .appname }} $(PWD)/main.go
run:
	go run $(PWD)/$(BIN_PATH)/{{ .appname }}
docker-build:
	GOOS=$(GOOS) go build -o $(PWD)/$(BIN_PATH)/{{ .appname }} $(PWD)/main.go

	docker build -f $(PWD)/Dockerfile -t image:tag $(PWD)

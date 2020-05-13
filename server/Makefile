# define global variable
SHELL:=/bin/bash
PROJ=ylops
ORG_PATH=casicloud.com/ylops/backend
REPO_PATH=$(ORG_PATH)/$(PROJ)
export PATH := $(PWD)/bin:$(PATH)
export GOBIN=$(PWD)/bin
# build number
BN=$(shell ./scripts/gen-bn.sh)
VERSION=$(shell cat VERSION).$(BN)

LD_FLAGS="-w -X $(REPO_PATH)/version.Version=$(VERSION)"

SRCS := $(shell find . -name '*.go'| grep -v vendor)

.PHONY: start build

NOW = $(shell date -u '+%Y%m%d%I%M%S')

APP = ${PROJ}

all: start

build:
	@go install -ldflags $(LD_FLAGS)  ./cmd/${APP}

start: 
	@echo ${PWD}
	@./bin/${APP} web -c ./configs/config.toml -m ./configs/model.conf --menu ./configs/menu.yaml

swagger:
	swag init --generalInfo ./internal/app/swagger.go --output ./internal/app/swagger

wire:
	wire gen ./internal/app/injector

test:
	@go test -v ./internal/app/test

clean:
	@rm -rf data bin ./internal/app/test/data ./cmd/${APP}/data

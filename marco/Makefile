# make file for alb project
# define global variable
SHELL:=/bin/bash
PROJ=marco
ORG_PATH=casicloud.com/ylops
MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
ROOT_DIR := $(patsubst %/,%,$(dir $(MAKEFILE_PATH)))
export PATH := $(ROOT_DIR)/bin:$(PATH)
export GOBIN=$(PWD)/bin
# build number
BN:=$(shell echo $$(date +"%Y%m%d%H%M").$$(git rev-parse --short HEAD).$$(git symbolic-ref --short -q HEAD))
VERSION=$(shell cat VERSION).$(BN)

LD_FLAGS="-w -X $(ROOT_DIR)/version.Version=$(VERSION)"

SRCS := $(shell find . -name '*.go'| grep -v vendor)

build: bin/polo

bin/polo: ${SRCS}
	@go install -v -ldflags $(LD_FLAGS) ${ROOT_DIR}/cmd/polo
	@cp -f ${ROOT_DIR}/cmd/polo/config.yaml  $(ROOT_DIR)/bin/

clean:
	@rm -rf bin/

.PHONY: version
version:
	@echo $(VERSION)

wire:
	wire gen ./internal/app/injector

swagger:
	swag init --generalInfo ./internal/app/swagger.go --output ./internal/app/swagger
test:
	@go test -v ./...

.PHONY: lint
lint: 
	@for file in $(SRCS); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

.PHONY: desc
desc:
	@echo root dir is $(ROOT_DIR)
	@echo version is $(VERSION)
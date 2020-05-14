# make file for alb project
# define global variable
SHELL:=/bin/bash
PROJ=backend
ORG_PATH=casicloud.com/ylops
REPO_PATH=$(ORG_PATH)/$(PROJ)
export PATH := $(PWD)/bin:$(PATH)
export GOBIN=$(PWD)/bin
# build number
BN=$(shell ./scripts/gen-bn.sh)
VERSION=$(shell cat VERSION).$(BN)

LD_FLAGS="-w -X $(REPO_PATH)/version.Version=$(VERSION)"

SRCS := $(shell find . -name '*.go'| grep -v vendor)

build: bin/ylops

bin/ylops:
	@go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/ylops

clean:
	@rm -rf bin/

.PHONY: version
version:
	@echo $(VERSION)

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
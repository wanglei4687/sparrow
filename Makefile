# https://clarkgrubb.com/makefile-style-guide
MAKEFLAGS += --warn-undefined-variables --no-builtin-rules
SHELL := /usr/bin/env bash
.SHELLFLAGS := -uo pipefail -c
.DEFAULT_GOAL := help
.DELETE_ON_ERROR:
.SUFFIXES:


SOURCES := $(shell find . -type f -name "*.go" -not -path "./bin/*" -not -path "./make/*")

GOBUILDPROCS ?=
TOOLS_BIN_DIR ?= $(shell pwd)/bin
GOLANGCILINTER_BINARY=$(TOOLS_BIN_DIR)/golangci-lint
TOOLING=$(GOLANGCILINTER_BINARY)
IMG_TAG ?= wanglei1995/sparrow:latest

include make/image.mk
include make/check.mk
include make/tools.mk

.PHONY: tools
tools: $(GOLANGCILINTER_BINARY)

.PHONY: clean-all
clean-all:
	rm -rf bin/

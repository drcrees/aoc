MAKEFLAGS += --silent
SHELL = bash

GOFMT_FILES?=$$(find . -name '*.go')

fmt: go-fmt

go-fmt:
	gofmt -w -s $(GOFMT_FILES)

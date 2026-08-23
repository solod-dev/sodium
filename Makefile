.PHONY: bind example help
.DEFAULT_GOAL := help

INC_DIR := $(shell pkg-config --variable=includedir libsodium 2>/dev/null)

ifeq ($(INC_DIR),)
  INC_DIR := /opt/homebrew/include
  FLAGS  := -I/opt/homebrew/include -L/opt/homebrew/lib
else
  FLAGS  := $(shell pkg-config --cflags-only-I --libs-only-L libsodium)
endif

CFLAGS ?= $(FLAGS)

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  bind                - Generate the bindings"
	@echo "  example name=<name> - Build an example program"

bind:
	@sobind -o libsodium/extern.go -pkg=libsodium -I $(INC_DIR) -scope=$(INC_DIR)/sodium -style=cap -strip=sodium_ -strip=crypto_ -rename=libsodium/rename.txt $(INC_DIR)/sodium.h

example:
	@CFLAGS="$(CFLAGS)" so build -check=sanitize -o ./build/$(name) ./example/$(name)

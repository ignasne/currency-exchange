MKFILE_PATH := $(lastword $(MAKEFILE_LIST))
CURRENT_DIR := $(patsubst %/,%,$(dir $(realpath $(MKFILE_PATH))))

vendor:
	go mod vendor

run-dev:
	go run "${CURRENT_DIR}/main.go" server

migrate:
	go run "${CURRENT_DIR}/main.go" migrate

.PHONY: vendor

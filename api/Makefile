MKFILE_PATH := $(lastword $(MAKEFILE_LIST))
CURRENT_DIR := $(patsubst %/,%,$(dir $(realpath $(MKFILE_PATH))))

vendor:
	go mod vendor

run-dev:
	go run "${CURRENT_DIR}/main.go" server

migrate:
	go run "${CURRENT_DIR}/main.go" migrate

test:
	go test -v -race ./router

build:
	docker build \
		-f .docker/Dockerfile \
		--force-rm \
		-t quotes-api \
		.

start:
	docker-compose up -d --force-recreate

stop:
	docker-compose down

.PHONY: vendor

.PHONY: run-dev build start stop

run-dev:
	yarn start

build:
	docker build \
		-f .docker/Dockerfile \
		--force-rm \
		-t currency-exchange-ui \
		.

start: build
	docker run -it --detach --publish=3000:80 --name=currency-exchange-ui currency-exchange-ui:latest

stop:
	docker stop currency-exchange-ui && docker rm currency-exchange-ui

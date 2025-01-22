GO_PROJECT_NAME := argus

build:
	docker compose build

up: build
	docker compose up

run: build
	docker compose up -d

stop:
	docker compose stop

.PHONY: build up run stop
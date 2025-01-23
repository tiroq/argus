GO_PROJECT_NAME := argus

build:
	docker compose build

up:
	docker compose up

run:
	docker compose up -d

stop:
	docker compose stop

.PHONY: build up run stop
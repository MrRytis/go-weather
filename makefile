DOCKER_GO_CONTAINER = docker-compose exec -T go

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

ssh:
	docker-compose exec go bash

migration_generate:
	$(DOCKER_GO_CONTAINER) | go run cmd/migration/migration_generate.go

migration_up:
	$(DOCKER_GO_CONTAINER) | go run cmd/migration/migration_up.go

migration_down:
	$(DOCKER_GO_CONTAINER) | go run cmd/migration/migration_down.go

run_air:
	$(DOCKER_GO_CONTAINER) | air

run:
	$(DOCKER_GO_CONTAINER) | go run main.go

swagger:
	$(DOCKER_GO_CONTAINER) | swag init

.PHONY: build up down ssh migration_generate migration_up migration_down run run_air
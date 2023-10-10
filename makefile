build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

ssh:
	docker-compose exec go bash

migration_generate:
	docker-compose exec go bash | go run cmd/migration/migration_generate.go

migration_up:
	docker-compose exec go bash | go run cmd/migration/migration_up.go

migration_down:
	docker-compose exec go bash | go run cmd/migration/migration_down.go

run_air:
	docker-compose exec go bash | air

run:
	docker-compose exec go bash | go run main.go

swagger:
	docker-compose exec go bash | swag init

.PHONY: build up down ssh migration_generate migration_up migration_down run run_air
build:
	docker-compose build

up:
	docker-compose up -d

migration_generate:
	go run cmd/migration/migration_generate.go

migration_up:
	go run cmd/migration/migration_up.go

migration_down:
	go run cmd/migration/migration_down.go

run:
	go run main.go

.PHONY: build up migration_generate migration_up migration_down
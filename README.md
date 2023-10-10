# Welcome to GO weather application

## Description

This is a simple weather application that uses the Meteo. lt and Open-Meteo API to get the weather data for a given city.

Currently support cities:
- Vilnius
- Kaunas
- Klaipeda

____

# Installation

1. Clone the repository
2. Run (first time):

```bash
make build
```

3. Migrate database:

```bash
make migration_up
```

4. Run the application:

```bash
make run
```

5. Open the browser and go to http://127.0.0.1:8080

___

# Make commands

Migrations:
- `make migration_generate` - Create new migration file
- `make migration_up` - Run migrations
- `make migration_down` - Rollback migrations

Run:
- `make run_air` - Run the application with hot reload
- `make run` - Run the application

Swagger:
- `make swagger` - Generate swagger documentation
- 
____

# Tech stack

- Go 
- Docker
- Docker-compose
- Postgres

# Main Libraries
- Gorilla Mux
- Gorm
- Cron/V3

____

## TODO
* Add tests



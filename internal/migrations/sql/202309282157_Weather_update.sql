-- +goose Up
ALTER TABLE weather_data
    ADD COLUMN "provider" varchar(255) NOT NULL DEFAULT '-';

-- +goose Down
ALTER TABLE weather_data
    DROP COLUMN "provider";
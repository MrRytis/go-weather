-- +goose Up
ALTER TABLE weather_data
    ALTER COLUMN pressure DROP NOT NULL,
    ALTER COLUMN humidity DROP NOT NULL,
    ALTER COLUMN clouds DROP NOT NULL;

-- +goose Down
ALTER TABLE weather_data
    ALTER COLUMN pressure SET NOT NULL,
    ALTER COLUMN humidity SET NOT NULL,
    ALTER COLUMN clouds SET NOT NULL;
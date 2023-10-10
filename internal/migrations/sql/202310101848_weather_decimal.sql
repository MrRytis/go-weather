-- +goose Up
ALTER TABLE weather_data
    ALTER COLUMN temp TYPE numeric(10,2),
    ALTER COLUMN feels_like TYPE numeric(10,2),
    ALTER COLUMN pressure TYPE numeric(10,2),
    ALTER COLUMN humidity TYPE integer,
    ALTER COLUMN wind_speed TYPE numeric(10,2),
    ALTER COLUMN wind_deg TYPE integer;


-- +goose Down
ALTER TABLE weather_data
    ALTER COLUMN temp TYPE FLOAT,
    ALTER COLUMN feels_like TYPE FLOAT,
    ALTER COLUMN pressure TYPE FLOAT,
    ALTER COLUMN humidity TYPE FLOAT,
    ALTER COLUMN wind_speed TYPE FLOAT,
    ALTER COLUMN wind_deg TYPE FLOAT;

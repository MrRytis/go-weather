-- +goose Up
CREATE TABLE weather_data
(
    id            SERIAL PRIMARY KEY,
    created_at    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    city          VARCHAR(255) NOT NULL,
    temp          FLOAT        NOT NULL,
    feels_like    FLOAT        NOT NULL,
    pressure      FLOAT        NOT NULL,
    humidity      FLOAT        NOT NULL,
    wind_speed    FLOAT        NOT NULL,
    wind_deg      FLOAT        NOT NULL,
    clouds        INTEGER      NOT NULL,
    weather       varchar(255) NOT NULL,
    precipitation INTEGER      NOT NULL,
    time          TIMESTAMP    NOT NULL
);

-- +goose Down
DROP TABLE weather_data;

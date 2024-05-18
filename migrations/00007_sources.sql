-- +goose Up
CREATE TYPE source_type AS ENUM ('telegram');
CREATE TABLE sources
(
    id         SERIAL PRIMARY KEY,
    title      TEXT UNIQUE NOT NULL,
    url        TEXT UNIQUE         NOT NULL,
    type       source_type         NOT NULL,
    country_id INTEGER             NOT NULL,

    FOREIGN KEY (country_id) REFERENCES countries (id)
);

-- +goose Down
DROP TABLE sources;
DROP TYPE source_type;
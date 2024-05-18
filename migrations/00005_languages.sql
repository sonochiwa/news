-- +goose Up
CREATE TABLE languages
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE languages;
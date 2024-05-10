-- +goose Up
CREATE TABLE languages
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE languages;
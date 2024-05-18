-- +goose Up
CREATE TABLE countries
(
    id   SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE countries;
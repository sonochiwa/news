-- +goose Up
CREATE TABLE countries
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(120) UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE countries;
-- +goose Up
CREATE TABLE categories
(
    id    SERIAL PRIMARY KEY,
    title TEXT,
    tag   TEXT
);

-- +goose Down
DROP TABLE categories;
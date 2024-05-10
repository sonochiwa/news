-- +goose Up
CREATE TABLE categories
(
    id SERIAL PRIMARY KEY
);

-- +goose Down
DROP TABLE categories;
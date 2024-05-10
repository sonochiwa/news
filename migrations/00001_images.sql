-- +goose Up
CREATE TABLE images
(
    id   SERIAL PRIMARY KEY,
    path TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE images;

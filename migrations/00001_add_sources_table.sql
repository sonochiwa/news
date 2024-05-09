-- +goose Up
CREATE TABLE sources
(
    id    int NOT NULL,
    title text,
    body  text,
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE sources;

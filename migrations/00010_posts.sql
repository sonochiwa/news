-- +goose Up
CREATE TABLE posts
(
    id         SERIAL PRIMARY KEY,
    image_id   INTEGER,
    source_id  INTEGER                 NOT NULL,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    deleted_at TIMESTAMP,

    FOREIGN KEY (image_id) REFERENCES images (id),
    FOREIGN KEY (source_id) REFERENCES sources (id)
);

-- +goose Down
DROP TABLE posts;
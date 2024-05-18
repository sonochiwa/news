-- +goose Up
ALTER TABLE posts
    ADD COLUMN title TEXT NOT NULL;

ALTER TABLE posts
    ADD COLUMN body TEXT NOT NULL;

-- +goose Down
ALTER TABLE posts
    DROP COLUMN title;

ALTER TABLE posts
    DROP COLUMN body;
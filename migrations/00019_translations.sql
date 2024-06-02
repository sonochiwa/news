-- +goose Up
CREATE TABLE translations
(
    id       SERIAL PRIMARY KEY,
    post_id  INTEGER REFERENCES posts (id) ON DELETE CASCADE,
    language VARCHAR(5),
    title    TEXT,
    body     TEXT,
    UNIQUE (post_id, language)
);

ALTER TABLE posts DROP COLUMN title;
ALTER TABLE posts DROP COLUMN body;

-- +goose Down
SELECT 'down SQL query';

-- +goose Up
ALTER table posts DROP column country_id;
ALTER TABLE posts ADD COLUMN country TEXT NULL;
ALTER TABLE posts ADD COLUMN category TEXT NULL;
ALTER TABLE posts DROP COLUMN deleted_at;
DROP TABLE languages;

-- +goose Down
SELECT 'down SQL query';

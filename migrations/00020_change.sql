-- +goose Up
ALTER TABLE posts DROP COLUMN country;
ALTER TABLE posts DROP COLUMN category;
ALTER TABLE translations ADD COLUMN country TEXT NULL;
ALTER TABLE translations ADD COLUMN category TEXT NULL;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

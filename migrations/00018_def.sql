-- +goose Up
ALTER TABLE users
ALTER COLUMN language SET DEFAULT 'ru';

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

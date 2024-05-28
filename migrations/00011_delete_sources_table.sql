-- +goose Up
ALTER TABLE posts DROP COLUMN source_id;
DROP TABLE sources;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

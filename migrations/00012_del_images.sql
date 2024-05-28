-- +goose Up
ALTER TABLE posts DROP column image_id;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

-- +goose Up
DROP TABLE countries;

-- +goose Down
SELECT 'down SQL query';

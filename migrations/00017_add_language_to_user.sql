-- +goose Up
alter table users
add column language text null;

-- +goose Down
SELECT 'down SQL query';

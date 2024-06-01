-- +goose Up
alter table posts add column country_tag text null;

-- +goose Down
SELECT 'down SQL query';

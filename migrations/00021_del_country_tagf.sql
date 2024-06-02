-- +goose Up
alter table posts drop column country_tag;

-- +goose Down
SELECT 'down SQL query';

-- +goose Up
ALTER TABLE posts
    ADD COLUMN country_id INTEGER NULL;

ALTER TABLE posts
    ADD CONSTRAINT fk_images
        FOREIGN KEY (country_id)
            REFERENCES countries (id);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

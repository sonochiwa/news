-- +goose Up
CREATE TABLE categories_languages
(
    title       VARCHAR(120),
    tag         VARCHAR(120),
    category_id INTEGER,
    language_id INTEGER,

    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE,
    FOREIGN KEY (language_id) REFERENCES languages (id) ON DELETE CASCADE,
    PRIMARY KEY (category_id, language_id)
);

-- +goose Down
DROP TABLE categories_languages;
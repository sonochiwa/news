-- +goose Up
CREATE TABLE posts_languages
(
    title       VARCHAR(120),
    body        TEXT NOT NULL,
    post_id     INTEGER,
    language_id INTEGER,

    FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
    FOREIGN KEY (language_id) REFERENCES languages (id) ON DELETE CASCADE,
    PRIMARY KEY (post_id, language_id)
);

-- +goose Down
DROP TABLE posts_languages;
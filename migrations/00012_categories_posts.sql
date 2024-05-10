-- +goose Up
CREATE TABLE categories_posts
(
    category_id INTEGER,
    post_id     INTEGER,

    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
    PRIMARY KEY (category_id, post_id)
);

-- +goose Down
DROP TABLE categories_posts;
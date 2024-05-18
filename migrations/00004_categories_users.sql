-- +goose Up
CREATE TABLE categories_users
(
    category_id INTEGER,
    user_id     INTEGER,

    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, category_id)
);

-- +goose Down
DROP TABLE categories_users;
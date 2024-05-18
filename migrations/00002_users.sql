-- +goose Up
CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    login         TEXT UNIQUE NOT NULL,
    password_hash TEXT        NOT NULL,
    image_id      INTEGER,
    is_admin      BOOLEAN     NOT NULL     DEFAULT FALSE,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    deleted_at    TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (image_id) REFERENCES images (id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE users;

-- +goose Up
CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(50) UNIQUE                     NOT NULL,
    email         VARCHAR(320) UNIQUE                    NOT NULL,
    password_hash VARCHAR(60)                            NOT NULL,
    image_id      INTEGER,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    deleted_at    TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (image_id) REFERENCES images (id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE users;

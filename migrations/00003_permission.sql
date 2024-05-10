-- +goose Up
CREATE TYPE permission_type AS ENUM ('super', 'watcher');
CREATE TABLE permissions
(
    user_id INTEGER PRIMARY KEY,
    type    permission_type NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- +goose Down
DROP TABLE permissions;
DROP TYPE permission_type;

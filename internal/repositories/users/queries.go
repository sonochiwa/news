package users

const (
	getAllUser = `
SELECT json_agg(row_to_json(users))
FROM (SELECT id, username, email, image_id, created_at, deleted_at
      FROM users AS u) users;
`
	getUserByID = `
SELECT (row_to_json(users))
FROM (SELECT id, username, email, image_id, created_at, deleted_at
      FROM users AS u
      WHERE id = $1) users;
`

	//user.Username, user.PasswordHash, user.Email, user.ImageId
	createUser = `
WITH new_user AS (
    INSERT INTO users (id, username, password_hash, email, image_id, created_at)
        VALUES (DEFAULT, $1, $2, $3, $4, DEFAULT)
        RETURNING *
)
SELECT row_to_json(new_user)
FROM new_user;
`

	getUserByEmail = `
SELECT (row_to_json(users))
FROM (
SELECT id, username, email, password_hash as password
FROM users AS u
WHERE email = $1
AND deleted_at IS NULL) users
`
)

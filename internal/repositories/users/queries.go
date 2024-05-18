package users

const (
	getAllUser = `
SELECT json_agg(row_to_json(users))
FROM (SELECT id, login, image_id, created_at, deleted_at
      FROM users AS u) users;
`
	getUserByID = `
SELECT (row_to_json(users))
FROM (SELECT id, login, image_id, created_at, deleted_at
      FROM users AS u
      WHERE id = $1) users;
`

	//user.Username, user.PasswordHash, user.Login, user.ImageId
	createUser = `
WITH new_user AS (
    INSERT INTO users (id, password_hash, login, image_id, created_at)
        VALUES (DEFAULT, $1, $2, $3, DEFAULT)
        RETURNING *
)
SELECT row_to_json(new_user)
FROM new_user;
`

	checkUser = `
SELECT (row_to_json(users))
FROM (
SELECT login, password_hash as password
FROM users AS u
WHERE login = $1
AND deleted_at IS NULL) users
`

	getUserByEmail = `
SELECT (row_to_json(users))
FROM (
SELECT id, login, password_hash as password
FROM users AS u
WHERE login = $1
AND deleted_at IS NULL) users
`
)

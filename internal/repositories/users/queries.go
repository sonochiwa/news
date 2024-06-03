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

	getUserByLogin = `
SELECT (row_to_json(users))
FROM (
	SELECT u.id, u.login, i.path as image_path, u.created_at, u.language, u.is_admin
	FROM users AS u
	LEFT JOIN images AS i ON u.image_id = i.id 
	WHERE login = $1
	AND deleted_at IS NULL
) users
`

	patchUserByLogin = `
UPDATE users
SET language = $2
WHERE login = $1
RETURNING id;
`
)

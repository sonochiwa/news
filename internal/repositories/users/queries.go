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
)

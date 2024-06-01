package posts

const (
	getAllPosts = `
SELECT COALESCE(json_agg(row_to_json(posts)), '[]'::json)
FROM (SELECT p.id         as id,
             p.created_at as created_at,
             p.title      as title,
             p.body       as body,
             p.country    as country,
             p.category   as category
      FROM posts AS p
	  WHERE ($1::text IS NULL OR $1 = '' OR p.title ILIKE concat('%', $1::text, '%'))
      AND ($2::text IS NULL OR $2 = '' OR p.category = $2::text)
      ) posts
`

	newPost = `
	INSERT INTO posts (id, title, body, category, country, created_at) 
	VALUES (DEFAULT, $1, $2, $3, $4, DEFAULT)
`
)

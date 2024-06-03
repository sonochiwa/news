package posts

const (
	getAllPosts = `
SELECT COALESCE(json_agg(row_to_json(posts)), '[]'::json)
FROM (SELECT p.id         as id,
             p.created_at as created_at,
             t.title      as title,
             t.body       as body,
             t.country    as country,
             t.category   as category,
             t.language   as language
      FROM posts AS p
      LEFT JOIN translations AS t on t.post_id = p.id
      WHERE ($4::text IS NULL OR $4 = '' OR t.language ILIKE $4::text)
        AND ($1::text IS NULL OR $1 = '' OR t.title ILIKE concat('%', $1::text, '%'))
        AND ($1::text IS NULL OR $1 = '' OR t.body ILIKE concat('%', $1::text, '%'))
        AND ($2::text IS NULL OR $2 = '' OR t.category ILIKE $2::text)
        AND ($3::text IS NULL OR $3 = '' OR t.country ILIKE $3::text)
      ORDER BY created_at DESC
) posts
`

	newPost = `
	INSERT INTO posts (id, created_at) 
	VALUES (DEFAULT, DEFAULT)
	RETURNING id
`
	newTranslation = `
	INSERT INTO translations (id, post_id, language, title, body, category, country) 
	VALUES (DEFAULT, $1, $2, $3, $4, $5, $6)
	RETURNING id
`
)

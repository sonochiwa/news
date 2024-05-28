package posts

const (
	getAllPosts = `
SELECT json_agg(row_to_json(posts))
FROM (SELECT p.id   as id,
             p.created_at as created_at,
             p.title 		as title,
             p.body 		as body,
      		 p.country 		as country,
      		 p.category 	as category
      FROM posts AS p) posts
`
)

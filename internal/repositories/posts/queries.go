package posts

const (
	getAllPosts = `
SELECT json_agg(row_to_json(posts))
FROM (SELECT p.id   as id,
       p.created_at as created_at,
       p.title 		as title,
       p.body 		as body,
       i.path       as image_path,
       s.title      as source_title,
       s.url        as source_url,
       s.type       as source_type,
       c.id         as country_id,
       c.name       as country_name
FROM posts AS p
         JOIN images as i ON p.image_id = i.id
         JOIN sources as s ON p.source_id = s.id
         JOIN countries as c on s.country_id = c.id
WHERE deleted_at is null

) posts`
)

package categories

const (
	getAllCategories = `
SELECT json_agg(row_to_json(categories))
FROM (SELECT DISTINCT p.category as title
      FROM posts AS p
     ) categories;
`
)

package categories

const (
	getAllCategories = `
SELECT COALESCE(json_agg(row_to_json(categories)), '[]'::json)
FROM (SELECT DISTINCT p.category as title
      FROM posts AS p
     ) categories;
`
)

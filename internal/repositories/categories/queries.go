package categories

const (
	getAllCategories = `
SELECT COALESCE(json_agg(row_to_json(categories)), '[]'::json)
FROM (SELECT DISTINCT t.category as title
      FROM translations AS t
      WHERE language = $1
     ) categories;
`
)

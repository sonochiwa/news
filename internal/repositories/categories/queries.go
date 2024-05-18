package categories

const (
	getAllCategories = `
SELECT json_agg(row_to_json(categories))
FROM (SELECT id, title, tag FROM categories AS c) categories;
`
)

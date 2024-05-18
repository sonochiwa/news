package languages

const (
	getAllLanguages = `
SELECT json_agg(row_to_json(languages))
FROM (SELECT id, name FROM languages AS l) languages;
`
)

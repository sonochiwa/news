package countries

const (
	getAllCountries = `
select COALESCE(json_agg(row_to_json(rows)), '[]'::json)
from (SELECT DISTINCT country as country_title
FROM translations
WHERE language = $1) rows

`
)

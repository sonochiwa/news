package countries

const (
	getAllCountries = `
select json_agg(rows) 
from (SELECT DISTINCT country as country_title, country_tag
FROM posts) rows

`
)

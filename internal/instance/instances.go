package instance

import "news/internal/common/postgres"

type Instances struct {
	Postgres postgres.Instance
}

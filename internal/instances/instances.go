package instances

import (
	"news/internal/instances/postgres"
)

type Instances struct {
	Postgres postgres.Instance
}

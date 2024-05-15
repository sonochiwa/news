package instances

import (
	"news/pkg/instances/postgres"
)

type Instances struct {
	Postgres postgres.Instance
}

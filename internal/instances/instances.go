package instances

import (
	"github.com/sonochiwa/news/internal/instances/postgres"
)

type Instances struct {
	Postgres postgres.Instance
}

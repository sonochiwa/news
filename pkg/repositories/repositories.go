package repositories

import (
	"news/pkg/instances"
	"news/pkg/repositories/users"
)

type Repositories struct {
	Users users.Repository
}

func New(db instances.Instances) Repositories {
	return Repositories{
		Users: users.New(db.Postgres),
	}
}

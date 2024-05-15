package repositories

import (
	"news/internal/instances"
	"news/internal/repositories/users"
)

type Repositories struct {
	Users users.Repository
}

func New(db instances.Instances) Repositories {
	return Repositories{
		Users: users.New(db.Postgres),
	}
}

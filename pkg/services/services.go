package services

import (
	"news/pkg/repositories"
	"news/pkg/services/users"
)

type Services struct {
	Users users.Services
}

func New(repository repositories.Repositories) Services {
	return Services{
		Users: users.New(repository),
	}
}

package services

import (
	"news/internal/repositories"
	"news/internal/services/users"
)

type Services struct {
	Users users.Services
}

func New(repository repositories.Repositories) Services {
	return Services{
		Users: users.New(repository),
	}
}

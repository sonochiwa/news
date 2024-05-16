package services

import (
	"github.com/sonochiwa/news/internal/repositories"
	"github.com/sonochiwa/news/internal/services/users"
)

type Services struct {
	Users users.Services
}

func New(repository repositories.Repositories) Services {
	return Services{
		Users: users.New(repository),
	}
}

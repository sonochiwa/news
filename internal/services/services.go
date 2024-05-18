package services

import (
	"github.com/sonochiwa/news/internal/repositories"
	"github.com/sonochiwa/news/internal/services/categories"
	"github.com/sonochiwa/news/internal/services/languages"
	"github.com/sonochiwa/news/internal/services/users"
)

type Services struct {
	Users      users.Services
	Languages  languages.Services
	Categories categories.Services
}

func New(repository repositories.Repositories) Services {
	return Services{
		Users:      users.New(repository),
		Languages:  languages.New(repository),
		Categories: categories.New(repository),
	}
}

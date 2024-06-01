package services

import (
	"github.com/sonochiwa/news/internal/repositories"
	"github.com/sonochiwa/news/internal/services/categories"
	"github.com/sonochiwa/news/internal/services/countries"
	"github.com/sonochiwa/news/internal/services/posts"
	"github.com/sonochiwa/news/internal/services/users"
)

type Services struct {
	Users      users.Services
	Categories categories.Services
	Posts      posts.Services
	Countries  countries.Services
}

func New(repository repositories.Repositories) Services {
	return Services{
		Users:      users.New(repository),
		Categories: categories.New(repository),
		Posts:      posts.New(repository),
		Countries:  countries.New(repository),
	}
}

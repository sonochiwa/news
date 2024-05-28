package repositories

import (
	"github.com/sonochiwa/news/internal/instances"
	"github.com/sonochiwa/news/internal/repositories/categories"
	"github.com/sonochiwa/news/internal/repositories/posts"
	"github.com/sonochiwa/news/internal/repositories/users"
)

type Repositories struct {
	Users      users.Repository
	Categories categories.Repository
	Posts      posts.Repository
}

func New(db instances.Instances) Repositories {
	return Repositories{
		Users:      users.New(db.Postgres),
		Categories: categories.New(db.Postgres),
		Posts:      posts.New(db.Postgres),
	}
}

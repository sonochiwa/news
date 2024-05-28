package posts

import (
	"github.com/sonochiwa/news/internal/instances/postgres"
	"github.com/sonochiwa/news/internal/models"

	"github.com/goccy/go-json"
)

type Postgres struct {
	db postgres.Instance
}

type Repository interface {
	GetAllPosts(filter string) (*[]models.Post, error)
}

func New(db postgres.Instance) Repository {
	return &Postgres{db: db}
}

func (p *Postgres) GetAllPosts(filter string) (result *[]models.Post, err error) {
	var bytes []byte

	if len(filter) > 0 {
		err = p.db.QueryRow(getAllPostsWithFilter, "%"+filter+"%").Scan(&bytes)
	} else {
		err = p.db.QueryRow(getAllPosts).Scan(&bytes)
	}

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

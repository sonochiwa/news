package posts

import (
	"fmt"

	"github.com/sonochiwa/news/internal/instances/postgres"
	"github.com/sonochiwa/news/internal/models"

	"github.com/goccy/go-json"
)

type Postgres struct {
	db postgres.Instance
}

type Repository interface {
	GetAllPosts(filter, category *string) (*[]models.Post, error)
}

func New(db postgres.Instance) Repository {
	return &Postgres{db: db}
}

func (p *Postgres) GetAllPosts(filter, category *string) (result *[]models.Post, err error) {
	var bytes []byte

	err = p.db.QueryRow(getAllPosts, *filter, *category).Scan(&bytes)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, fmt.Errorf("repository json.Unmarshal: %w", err)
	}

	return result, nil
}

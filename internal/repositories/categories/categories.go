package categories

import (
	"github.com/sonochiwa/news/internal/instances/postgres"
	"github.com/sonochiwa/news/internal/models"

	"github.com/goccy/go-json"
)

type Postgres struct {
	db postgres.Instance
}

type Repository interface {
	GetAllCategories() (*[]models.Categories, error)
}

func New(db postgres.Instance) Repository {
	return &Postgres{db: db}
}

func (p *Postgres) GetAllCategories() (result *[]models.Categories, err error) {
	var bytes []byte

	err = p.db.QueryRow(getAllCategories).Scan(&bytes)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

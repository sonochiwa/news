package languages

import (
	"github.com/sonochiwa/news/internal/instances/postgres"
	"github.com/sonochiwa/news/internal/models"

	"github.com/goccy/go-json"
)

type Postgres struct {
	db postgres.Instance
}

type Repository interface {
	GetAllLanguages() (*[]models.Language, error)
}

func New(db postgres.Instance) Repository {
	return &Postgres{db: db}
}

func (p *Postgres) GetAllLanguages() (result *[]models.Language, err error) {
	var bytes []byte

	err = p.db.QueryRow(getAllLanguages).Scan(&bytes)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

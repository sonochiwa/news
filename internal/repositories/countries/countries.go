package countries

import (
	"github.com/sonochiwa/news/internal/instances/postgres"
	"github.com/sonochiwa/news/internal/models"

	"github.com/goccy/go-json"
)

type Postgres struct {
	db postgres.Instance
}

type Repository interface {
	GetAllCountries() (*[]models.Country, error)
}

func New(db postgres.Instance) Repository {
	return &Postgres{db: db}
}

func (p *Postgres) GetAllCountries() (result *[]models.Country, err error) {
	var bytes []byte

	err = p.db.QueryRow(getAllCountries).Scan(&bytes)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

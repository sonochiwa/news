package users

import (
	"news/internal/instances/postgres"
	"news/internal/models"

	"github.com/goccy/go-json"
)

type Postgres struct {
	db postgres.Instance
}

type Repository interface {
	GetAllUsers() (*[]models.User, error)
	GetUserByID(id int64) (*models.User, error)
}

func New(db postgres.Instance) Repository {
	return &Postgres{db: db}
}

func (p *Postgres) GetAllUsers() (result *[]models.User, err error) {
	var bytes []byte

	err = p.db.QueryRow(getAllUser).Scan(&bytes)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Postgres) GetUserByID(id int64) (result *models.User, err error) {
	var bytes []byte

	err = p.db.QueryRow(getUserByID, id).Scan(&bytes)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

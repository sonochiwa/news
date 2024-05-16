package users

import (
	"errors"

	"github.com/sonochiwa/news/internal/instances/postgres"
	"github.com/sonochiwa/news/internal/models"

	"github.com/goccy/go-json"
)

type Postgres struct {
	db postgres.Instance
}

type Repository interface {
	GetAllUsers() (*[]models.User, error)
	GetUserByID(id int64) (*models.User, error)
	CreateUser(user *models.User) (result *models.User, err error)
	GetUserByEmail(email string) (*models.User, error)
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

func (p *Postgres) CreateUser(user *models.User) (result *models.User, err error) {
	var bytes []byte

	err = p.db.QueryRow(createUser, user.Username, user.PasswordHash, user.Email, user.ImageId).Scan(&bytes)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Postgres) GetUserByEmail(email string) (result *models.User, err error) {
	var bytes []byte

	err = p.db.QueryRow(getUserByEmail, email).Scan(&bytes)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Email) == 0 {
		return nil, errors.New("user not found")
	}

	return result, nil
}

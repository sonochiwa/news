package users

import (
	"errors"
	"fmt"

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
	CheckUser(email string) (*models.SignInUser, error)
	GetUserByLogin(email string) (*models.UserMe, error)
	UpdateUserPhoto(userID int, imagePath string) (err error)
}

func New(db postgres.Instance) Repository {
	return &Postgres{db: db}
}

func (p *Postgres) UpdateUserPhoto(userID int, imagePath string) (err error) {
	tx, err := p.db.DB().Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var imageID int
	err = tx.QueryRow("INSERT INTO images (path) VALUES ($1) RETURNING id", imagePath).Scan(&imageID)
	if err != nil {
		return fmt.Errorf("failed to insert image: %v", err)
	}

	_, err = tx.Exec("UPDATE users SET image_id = $1 WHERE id = $2", imageID, userID)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return err
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

	err = p.db.QueryRow(createUser, user.PasswordHash, user.Login, user.ImageId).Scan(&bytes)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Postgres) CheckUser(email string) (result *models.SignInUser, err error) {
	var bytes []byte

	err = p.db.QueryRow(checkUser, email).Scan(&bytes)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Login) == 0 {
		return nil, errors.New("user not found")
	}

	return result, nil
}

func (p *Postgres) GetUserByLogin(email string) (result *models.UserMe, err error) {
	var bytes []byte

	err = p.db.QueryRow(getUserByLogin, email).Scan(&bytes)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Login) == 0 {
		return nil, errors.New("user not found")
	}

	return result, nil
}

package users

import (
	"github.com/sonochiwa/news/internal/models"
	"github.com/sonochiwa/news/internal/repositories"
	"github.com/sonochiwa/news/internal/utils"
)

type Service struct {
	repository repositories.Repositories
}

type Services interface {
	GetAllUsers() (*[]models.User, error)
	GetUserByID(id int64) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	CheckUser(email string) (result *models.SignInUser, err error)
	GetUserByEmail(email string) (result *models.User, err error)
}

func New(repository repositories.Repositories) Services {
	return &Service{repository: repository}
}

func (s *Service) GetAllUsers() (*[]models.User, error) {
	users, _ := s.repository.Users.GetAllUsers()
	return users, nil
}

func (s *Service) GetUserByID(id int64) (*models.User, error) {
	user, _ := s.repository.Users.GetUserByID(id)
	return user, nil
}

func (s *Service) CreateUser(user *models.User) (*models.User, error) {
	user.PasswordHash, _ = utils.HashPassword(user.PasswordHash)
	user, err := s.repository.Users.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) CheckUser(email string) (result *models.SignInUser, err error) {
	result, err = s.repository.Users.CheckUser(email)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) GetUserByEmail(email string) (result *models.User, err error) {
	result, err = s.repository.Users.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return result, nil
}

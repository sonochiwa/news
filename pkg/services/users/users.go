package users

import (
	"news/pkg/models"
	"news/pkg/repositories"
)

type Service struct {
	repository repositories.Repositories
}

type Services interface {
	GetAllUsers() (*[]models.User, error)
	GetUserByID(id int64) (*models.User, error)
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

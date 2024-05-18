package languages

import (
	"github.com/sonochiwa/news/internal/models"
	"github.com/sonochiwa/news/internal/repositories"
)

type Service struct {
	repository repositories.Repositories
}

type Services interface {
	GetAllLanguages() (*[]models.Language, error)
}

func New(repository repositories.Repositories) Services {
	return &Service{repository: repository}
}

func (s *Service) GetAllLanguages() (*[]models.Language, error) {
	users, _ := s.repository.Languages.GetAllLanguages()
	return users, nil
}

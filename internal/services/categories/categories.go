package categories

import (
	"github.com/sonochiwa/news/internal/models"
	"github.com/sonochiwa/news/internal/repositories"
)

type Service struct {
	repository repositories.Repositories
}

type Services interface {
	GetAllCategories(language string) (*[]models.Category, error)
}

func New(repository repositories.Repositories) Services {
	return &Service{repository: repository}
}

func (s *Service) GetAllCategories(language string) (*[]models.Category, error) {
	categories, _ := s.repository.Categories.GetAllCategories(language)
	return categories, nil
}

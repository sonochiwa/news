package countries

import (
	"github.com/sonochiwa/news/internal/models"
	"github.com/sonochiwa/news/internal/repositories"
)

type Service struct {
	repository repositories.Repositories
}

type Services interface {
	GetAllCountries(language string) (*[]models.Country, error)
}

func New(repository repositories.Repositories) Services {
	return &Service{repository: repository}
}

func (s *Service) GetAllCountries(language string) (*[]models.Country, error) {
	countries, _ := s.repository.Countries.GetAllCountries(language)

	return countries, nil
}

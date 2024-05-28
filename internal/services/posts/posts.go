package posts

import (
	"github.com/sonochiwa/news/internal/models"
	"github.com/sonochiwa/news/internal/repositories"
)

type Service struct {
	repository repositories.Repositories
}

type Services interface {
	GetAllPosts(filter string) (*[]models.Post, error)
}

func New(repository repositories.Repositories) Services {
	return &Service{repository: repository}
}

func (s *Service) GetAllPosts(title string) (*[]models.Post, error) {
	posts, _ := s.repository.Posts.GetAllPosts(title)
	return posts, nil
}

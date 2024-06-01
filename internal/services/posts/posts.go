package posts

import (
	"fmt"

	"github.com/sonochiwa/news/internal/models"
	"github.com/sonochiwa/news/internal/repositories"
)

type Service struct {
	repository repositories.Repositories
}

type Services interface {
	GetAllPosts(filter, category *string) (*[]models.Post, error)
	NewPost(input models.NewPost) error
}

func New(repository repositories.Repositories) Services {
	return &Service{repository: repository}
}

func (s *Service) GetAllPosts(filter, category *string) (*[]models.Post, error) {
	posts, err := s.repository.Posts.GetAllPosts(filter, category)
	if err != nil {
		return nil, fmt.Errorf("service: %w", err)
	}

	return posts, nil
}

func (s *Service) NewPost(input models.NewPost) error {
	err := s.repository.Posts.NewPost(input)
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return nil
}

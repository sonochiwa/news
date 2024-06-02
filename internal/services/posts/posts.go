package posts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sonochiwa/news/configs"
	"github.com/sonochiwa/news/internal/models"
	"github.com/sonochiwa/news/internal/repositories"
)

type Service struct {
	repository repositories.Repositories
}

type Services interface {
	GetAllPosts(filter, category, country, language *string) (*[]models.Post, error)
	NewPost(input models.NewPost) error
	DeletePost(id int) error
}

func New(repository repositories.Repositories) Services {
	return &Service{repository: repository}
}

func (s *Service) GetAllPosts(filter, category, country, language *string) (*[]models.Post, error) {
	posts, err := s.repository.Posts.GetAllPosts(filter, category, country, language)
	if err != nil {
		return nil, fmt.Errorf("service: %w", err)
	}

	return posts, nil
}

type translationResponse struct {
	TranslatedText []string `json:"translatedText"`
}

func (s *Service) NewPost(input models.NewPost) error {
	var languages = []string{"ru", "en", "de", "pt", "zh"}

	postID, err := s.repository.Posts.NewPost()
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	for _, v := range languages {
		requestBody := map[string]interface{}{
			"q":      []string{input.Title, input.Body, input.Category, input.Country},
			"source": "auto",
			"target": v,
			"format": "text",
		}

		jsonValue, err := json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("service json.Marshal: %w", err)
		}

		resp, err := http.Post(
			configs.GlobalConfig.Other.LibreDomain,
			"application/json",
			bytes.NewBuffer(jsonValue),
		)
		if err != nil {
			return fmt.Errorf("error http.Post: %w", err)
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %w", err)
		}

		var responseObject *translationResponse
		err = json.Unmarshal(bodyBytes, &responseObject)
		if err != nil {
			return fmt.Errorf("error unmarshalling JSON: %w", err)
		}

		t := responseObject.TranslatedText

		input.Title = t[0]
		input.Body = t[1]
		input.Category = t[2]
		input.Country = t[3]

		err = s.repository.Posts.NewTranslation(postID, v, input)
		if err != nil {
			return fmt.Errorf("service: %w", err)
		}
	}

	return nil
}

func (s *Service) DeletePost(id int) error {
	err := s.repository.Posts.DeletePost(id)
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return nil
}

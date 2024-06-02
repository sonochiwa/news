package posts

import (
	"fmt"
	"strings"

	"github.com/sonochiwa/news/internal/instances/postgres"
	"github.com/sonochiwa/news/internal/models"

	"github.com/goccy/go-json"
)

type Postgres struct {
	db postgres.Instance
}

type Repository interface {
	GetAllPosts(filter, category, country, language *string) (*[]models.Post, error)
	NewPost(input models.NewPost) (postID int, err error)
	NewTranslation(postID int, language string, input models.NewPost) (err error)
	DeletePost(id int) (err error)
}

func New(db postgres.Instance) Repository {
	return &Postgres{db: db}
}

func (p *Postgres) GetAllPosts(filter, category, country, language *string) (result *[]models.Post, err error) {
	var bytes []byte

	err = p.db.QueryRow(getAllPosts, *filter, *category, *country, *language).Scan(&bytes)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, fmt.Errorf("repository json.Unmarshal: %w", err)
	}

	return result, nil
}

func (p *Postgres) NewPost(input models.NewPost) (postID int, err error) {
	var response int

	//err = p.db.QueryRow(newPost, input.Title, input.Body, strings.ToLower(input.Category), input.Country, strings.ToLower(input.CountryTag)).Scan(&bytes)
	err = p.db.QueryRow(newPost, strings.ToLower(input.CountryTag)).Scan(&response)
	if err != nil {
		return 0, err
	}

	return response, nil
}

func (p *Postgres) NewTranslation(postID int, language string, input models.NewPost) (err error) {
	var bytes []byte

	err = p.db.QueryRow(newTranslation,
		postID, language, input.Title, input.Body, strings.ToLower(input.Category), input.Country,
	).Scan(&bytes)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeletePost(id int) error {
	var bytes []byte

	err := p.db.QueryRow("delete from posts where id = $1 returning id", id).Scan(&bytes)
	if err != nil {
		return err
	}

	return nil
}

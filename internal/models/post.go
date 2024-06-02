package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Country   string    `json:"country"`
	Category  string    `json:"category"`
	Language  string    `json:"language"`
}

type NewPost struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	Country    string `json:"country"`
	CountryTag string `json:"country_tag"`
	Category   string `json:"category"`
}

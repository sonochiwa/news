package models

import "time"

type Post struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	SourceTitle string    `json:"source_title"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	SourceUrl   string    `json:"source_url"`
	SourceType  string    `json:"source_type"`
	CountryID   int       `json:"country_id"`
	CountryName string    `json:"country_name"`
}

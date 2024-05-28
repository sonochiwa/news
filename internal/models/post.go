package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Country   string    `json:"country"`
	Category  string    `json:"category"`
}

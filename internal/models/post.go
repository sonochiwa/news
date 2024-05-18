package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	ImageID   int       `json:"image_id"`
	SourceID  int       `json:"source_id"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

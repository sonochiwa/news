package models

import "time"

type User struct {
	ID           int64      `json:"id"`
	Login        string     `json:"login"`
	PasswordHash string     `json:"password,omitempty"`
	ImageId      *int64     `json:"image_id"`
	CreatedAt    time.Time  `json:"created_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type UserMe struct {
	ID        int64     `json:"id"`
	Login     string    `json:"login"`
	ImagePath *string   `json:"image_path"`
	CreatedAt time.Time `json:"created_at"`
}

type SignInUser struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

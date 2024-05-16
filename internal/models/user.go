package models

import "time"

type User struct {
	ID           int64      `json:"id"`
	Username     string     `json:"username""`
	Email        string     `json:"email"`
	PasswordHash string     `json:"password,omitempty""`
	ImageId      *int64     `json:"image_id"`
	CreatedAt    time.Time  `json:"created_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type SignInUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

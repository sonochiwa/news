package models

type Categories struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Tag   *string `json:"tag"`
}

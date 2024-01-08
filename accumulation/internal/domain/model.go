package domain

type Point struct {
	ID         string  `json:"id"`
	User       string  `json:"user"`
	Total      float32 `json:"total"`
	CreateDate string
}

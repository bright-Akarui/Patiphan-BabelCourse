package models

// import "gorm.io/gorm"

type Article struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Image string `json:"image"`
	// gorm.Model
}

package models

import "gorm.io/gorm"

type Article struct {
	// ID    uint   `json:"id"`
	// Title string `json:"title"`
	// Body  string `json:"body"`
	// Image string `json:"image"`
	gorm.Model
	Title       string `gorm:"unique;not null"`
	Excerpt     string `gorm:"not null"`
	Body        string `gorm:"not null"`
	Image       string `gorm:"not null"`
	ID_Category int
	Category    Category `gorm:"foreignKey:ID_Category"`
}

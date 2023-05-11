package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title      string
	Body       string
	ImageUrl   string
	UserClicks []UserClicks `gorm:"foreignKey:PostID"`
}

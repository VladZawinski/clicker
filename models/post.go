package models

import "gorm.io/gorm"

type Post struct {
	Title       string
	Body        string
	ImageUrl    string
	Url         string
	ContentType string
	UserClicks  []UserClicks `gorm:"foreignKey:PostID"`
	gorm.Model
}

type PostWithCount struct {
	Post
	UserClicksCount int
}

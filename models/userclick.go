package models

import "gorm.io/gorm"

type UserClicks struct {
	gorm.Model
	UserID uint
	PostID uint
	Count  uint
}

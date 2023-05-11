package models

import (
	"time"

	"gorm.io/gorm"
)

type UserClicks struct {
	UserID    uint `gorm:"not null"`
	PostID    uint `gorm:"not null"`
	Ip        *string
	User      User `gorm:"foreignKey:UserID"`
	Post      Post `gorm:"foreignKey:PostID"`
	ClickedAt time.Time
	gorm.Model
}

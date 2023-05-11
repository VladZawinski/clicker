package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string
	Phone        string
	Password     string
	ClickedCount uint
	Role         string
	UserClicks   []UserClicks `gorm:"foreignKey:UserID"`
}

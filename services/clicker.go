package services

import "gorm.io/gorm"

type ClickerService struct {
	Post PostService
	User UserService
}

func NewClickerService(db *gorm.DB) ClickerService {
	return ClickerService{}
}

package services

import (
	"clicker/models"

	"gorm.io/gorm"
)

type UserService interface {
	RegisterUser(user *models.User) error
	FindUserByPhone(phone string) (*models.User, error)
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

type userService struct {
	db *gorm.DB
}

func (s *userService) RegisterUser(user *models.User) error {
	return nil
}

func (s *userService) FindUserByPhone(phone string) (*models.User, error) {
	return nil, nil
}

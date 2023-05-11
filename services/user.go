package services

import (
	"clicker/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService interface {
	RegisterUser(user *models.User) error
	FindUserByPhone(phone string) (*models.User, error)
	FindAllUser() ([]models.User, error)
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

type userService struct {
	db *gorm.DB
}

func (s *userService) RegisterUser(user *models.User) error {
	result := s.db.Create(user)
	if result.RowsAffected != 0 {
		return nil
	}
	return fiber.ErrBadRequest
}

func (s *userService) FindUserByPhone(phone string) (*models.User, error) {
	var result *models.User
	err := s.db.Where("Phone = ?", phone).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) FindAllUser() ([]models.User, error) {
	var users []models.User
	result := s.db.Find(&users)
	fmt.Println(result)
	return users, nil
}

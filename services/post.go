package services

import "gorm.io/gorm"

type PostService interface {
}

type postService struct {
	db *gorm.DB
}

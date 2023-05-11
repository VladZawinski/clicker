package services

import (
	"clicker/models"

	"gorm.io/gorm"
)

type PostService interface {
	CreatePost(post *models.Post) error
	UpdatePost(post *models.Post) error
	DeletePost(id uint) error
	GetAllPosts() ([]models.Post, error)
	GetPostByID(id uint) (*models.Post, error)
}

func NewPostService(db *gorm.DB) PostService {
	return &postService{
		db: db,
	}
}

type postService struct {
	db *gorm.DB
}

func (s *postService) CreatePost(post *models.Post) error {
	return s.db.Create(post).Error
}

func (s *postService) UpdatePost(post *models.Post) error {
	return s.db.Save(post).Error
}

func (s *postService) DeletePost(id uint) error {
	return s.db.Delete(&models.Post{}, id).Error
}

func (s *postService) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := s.db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *postService) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	err := s.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

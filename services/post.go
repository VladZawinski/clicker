package services

import (
	"clicker/dto"
	"clicker/models"
	"time"

	"gorm.io/gorm"
)

type PostService interface {
	CreatePost(post *models.Post) error
	UpdatePost(post *models.Post) error
	DeletePost(id uint) error
	GetAllPosts() ([]models.PostWithCount, error)
	GetPostByID(id int) (*models.Post, error)
	MarkPostAsClicked(userId int, id int) error
	GetAllUserClicks() ([]models.UserClicks, error)
	GetClickedUsersById(id int) ([]dto.ClickedPostUser, error)
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

func (s *postService) GetAllPosts() ([]models.PostWithCount, error) {
	var posts []models.Post
	err := s.db.Model(&models.Post{}).
		Preload("UserClicks.User").
		Preload("UserClicks").
		Find(&posts).Error
	if err != nil {
		return nil, err
	}
	postsWithCount := make([]models.PostWithCount, len(posts))

	for i := range posts {
		postsWithCount[i].Post = posts[i]
		postsWithCount[i].UserClicksCount = len(posts[i].UserClicks)
	}
	return postsWithCount, nil
}

func (s *postService) GetPostByID(id int) (*models.Post, error) {
	var post models.Post
	err := s.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *postService) MarkPostAsClicked(userId int, id int) error {
	var user models.User
	var post models.Post
	s.db.First(&user, userId)
	s.db.First(&post, id)
	return s.db.Create(&models.UserClicks{
		UserID:    uint(userId),
		PostID:    uint(id),
		ClickedAt: time.Now(),
		User:      user,
		Post:      post,
	}).Error
}

func (s *postService) GetAllUserClicks() ([]models.UserClicks, error) {
	var clicks []models.UserClicks
	err := s.db.Preload("User").Preload("Post").Find(&clicks).Error
	if err != nil {
		return nil, err
	}
	return clicks, nil
}

func (s *postService) GetClickedUsersById(id int) ([]dto.ClickedPostUser, error) {
	var post models.Post
	var userCount []dto.ClickedPostUser
	userClicksCount := make(map[uint]int)
	err := s.db.Preload("UserClicks.User").Find(&post).Error
	if err != nil {
		return nil, err
	}
	for _, uc := range post.UserClicks {
		userClicksCount[uc.UserID]++
	}
	userMap := make(map[uint]*models.User)
	for _, uc := range post.UserClicks {
		user := uc.User
		if _, ok := userMap[user.ID]; !ok {
			userMap[user.ID] = &user
		}
	}

	for userID, user := range userMap {
		count := userClicksCount[userID]
		userCount = append(userCount, dto.ClickedPostUser{Id: int(user.ID), User: dto.User{Id: user.ID, Name: user.Name, Phone: user.Phone}, Count: count})
	}
	return userCount, nil
}

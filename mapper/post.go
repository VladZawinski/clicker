package mapper

import (
	"clicker/dto"
	"clicker/models"
)

func MapPostToDto(p *models.Post) *dto.Post {
	return &dto.Post{
		Id:          p.ID,
		Title:       p.Title,
		Body:        p.Body,
		ImageUrl:    p.ImageUrl,
		Url:         p.Url,
		CreateAt:    p.CreatedAt,
		ContentType: p.ContentType,
	}
}

func MapPostsToPostDtos(posts []models.Post) []*dto.Post {
	dtos := make([]*dto.Post, len(posts))
	for i, p := range posts {
		dtos[i] = MapPostToDto(&p)
	}
	return dtos
}

package dto

import "time"

type Post struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	Url         string    `json:"url"`
	ImageUrl    string    `json:"imageUrl"`
	ContentType string    `json:"contentType"`
	CreateAt    time.Time `json:"createdAt"`
}

type CreatePost struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	Url         string `json:"url"`
	ImageUrl    string `json:"imageUrl"`
	ContentType string `json:"contentType"`
}

type UpdatePost struct {
	Title    *string `json:"title"`
	Body     *string `json:"body"`
	Url      *string `json:"url"`
	ImageUrl *string `json:"imageUrl"`
}

package handlers

import (
	"clicker/services"

	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	service services.ClickerService
}

func NewPostHandler(s services.ClickerService) PostHandler {
	return PostHandler{service: s}
}

func (h *PostHandler) GetAllPost(c *fiber.Ctx) error {
	return nil
}

func (h *PostHandler) GetPostById(c *fiber.Ctx) error {
	return nil
}

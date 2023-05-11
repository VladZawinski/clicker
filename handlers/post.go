package handlers

import (
	"clicker/mapper"
	"clicker/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	service services.ClickerService
}

func NewPostHandler(s services.ClickerService) PostHandler {
	return PostHandler{service: s}
}

func (h *PostHandler) GetAllPost(c *fiber.Ctx) error {
	result, err := h.service.Post.GetAllPosts()
	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "Something went wrong")
	}
	return c.JSON(mapper.MapPostsToPostDtos(result))
}

func (h *PostHandler) GetPostById(c *fiber.Ctx) error {
	param := c.Params("name")
	id, _ := strconv.Atoi(param)
	result, err := h.service.Post.GetPostByID(id)
	if err != nil {
		return err
	}
	return c.JSON(mapper.MapPostToDto(result))
}

func (h *PostHandler) MarkAsClicked(c *fiber.Ctx) error {
	return nil
}

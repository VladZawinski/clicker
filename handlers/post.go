package handlers

import (
	"clicker/mapper"
	"clicker/middlewares"
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
	result, err := h.service.Post.GetAllPosts()
	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "Something went wrong")
	}
	return c.JSON(mapper.MapPostsToPostDtos(result))
}

func (h *PostHandler) GetPostById(c *fiber.Ctx) error {
	id, _ := middlewares.ParseIdParam(c)
	result, err := h.service.Post.GetPostByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}
	return c.JSON(mapper.MapPostToDto(result))
}

func (h *PostHandler) MarkAsClicked(c *fiber.Ctx) error {
	id, err := middlewares.ExtractUser(c)
	postId, _ := middlewares.ParseIdParam(c)
	if err != nil {
		return err
	}
	errr := h.service.Post.MarkPostAsClicked(int(id), postId)
	if errr != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}

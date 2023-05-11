package handlers

import (
	"clicker/dto"
	"clicker/models"
	"clicker/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type PortalHandler struct {
	service services.ClickerService
}

func NewPortalHandler(s services.ClickerService) PortalHandler {
	return PortalHandler{service: s}
}

func (h *PortalHandler) Login(c *fiber.Ctx) error {
	return nil
}

func (h *PortalHandler) CreatePost(c *fiber.Ctx) error {
	body := new(dto.CreatePost)
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fields are required")
	}
	err := h.service.Post.CreatePost(&models.Post{
		Title:       body.Title,
		Body:        body.Body,
		ImageUrl:    body.ImageUrl,
		Url:         body.Url,
		ContentType: body.ContentType,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	return nil
}

func (h *PortalHandler) DeletePost(c *fiber.Ctx) error {
	return nil
}

func (h *PortalHandler) FindPostDetail(c *fiber.Ctx) error {
	return nil
}

func (h *PortalHandler) FindAllUser(c *fiber.Ctx) error {
	result, err := h.service.User.FindAllUser()
	fmt.Println(result)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	return c.JSON(result)
}

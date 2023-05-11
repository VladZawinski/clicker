package handlers

import (
	"clicker/services"

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
	return nil
}

func (h *PortalHandler) DeletePost(c *fiber.Ctx) error {
	return nil
}

func (h *PortalHandler) FindPostDetail(c *fiber.Ctx) error {
	return nil
}

func (h *PortalHandler) FindAllUser(c *fiber.Ctx) error {
	return nil
}

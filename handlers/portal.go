package handlers

import (
	"clicker/dto"
	"clicker/mapper"
	"clicker/middlewares"
	"clicker/models"
	"clicker/services"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type PortalHandler struct {
	service services.ClickerService
}

func NewPortalHandler(s services.ClickerService) PortalHandler {
	return PortalHandler{service: s}
}

func (h *PortalHandler) Login(c *fiber.Ctx) error {
	body := new(dto.Login)
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fields are required")
	}
	u, _ := h.service.User.FindUserByPhone(body.Phone)
	if u == nil {
		return fiber.NewError(fiber.StatusForbidden)
	}
	if u.Password != body.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}
	if u.Role != middlewares.AdminRole {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "error", "message": "Forbidden", "data": nil})
	}
	claims := jwt.MapClaims{
		"username": u.Phone,
		"userId":   u.ID,
		"role":     u.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
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
	param := c.Params("id")
	id, _ := strconv.Atoi(param)
	err := h.service.Post.DeletePost(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return nil
}

func (h *PortalHandler) FindPostDetail(c *fiber.Ctx) error {
	param := c.Params("id")
	id, _ := strconv.Atoi(param)
	result, err := h.service.Post.GetPostByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}
	return c.JSON(mapper.MapPostToDto(result))
}

func (h *PortalHandler) FindAllUser(c *fiber.Ctx) error {
	result, err := h.service.User.FindAllUser()
	fmt.Println(result)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	return c.JSON(result)
}

func (h *PortalHandler) GetAllUserClicks(c *fiber.Ctx) error {
	result, err := h.service.Post.GetAllUserClicks()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	return c.JSON(result)
}

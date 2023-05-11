package handlers

import (
	"clicker/dto"
	"clicker/middlewares"
	"clicker/models"
	"clicker/services"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	service services.ClickerService
}

func NewAuthHandler(s services.ClickerService) AuthHandler {
	return AuthHandler{service: s}
}

func (h AuthHandler) Login(c *fiber.Ctx) error {
	body := new(dto.Login)
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fields are required")
	}
	u, err := h.service.User.FindUserByPhone(body.Phone)
	fmt.Println(err)
	if u == nil {
		return fiber.NewError(fiber.StatusForbidden)
	}
	if u.Password != body.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}
	claims := jwt.MapClaims{
		"username": u.Phone,
		"userId":   u.ID,
		"role":     u.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}

func (h AuthHandler) Register(c *fiber.Ctx) error {
	body := new(dto.RegisterUser)
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fields are required")
	}
	user, err := h.service.User.FindUserByPhone(body.Phone)
	if err != nil && user != nil {
		return fiber.NewError(fiber.StatusConflict)
	}
	h.service.User.RegisterUser(&models.User{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: body.Password,
		Role:     middlewares.UserRole,
	})
	return c.SendStatus(fiber.StatusCreated)
}

package middlewares

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParseIdParam(c *fiber.Ctx) (int, error) {
	param := c.Params("id")
	return strconv.Atoi(param)
}

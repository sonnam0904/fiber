package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func UserMiddleware(c *fiber.Ctx) error {
	fmt.Println("Checked User", c)

	return c.Next()
}
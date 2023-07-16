package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func HelloWorld() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		name := ctx.Query("name", "user")

		return ctx.Status(fiber.StatusOK).JSON(fmt.Sprintf("Hello, %s!", name))
	}
}

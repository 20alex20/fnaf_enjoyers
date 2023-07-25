package handler

import (
	"github.com/fnaf-enjoyers/user/config"
	"github.com/gofiber/fiber/v2"
)

func LogOut() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := config.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
		}

		err = session.Destroy()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON("success")
	}
}

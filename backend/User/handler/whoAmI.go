package handler

import (
	"github.com/fnaf-enjoyers/user/config"
	"github.com/gofiber/fiber/v2"
)

func WhoAmI() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := config.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
		}

		name := session.Get("name")
		id := session.ID()
		keys := session.Keys()
		fresh := session.Fresh()

		json := struct {
			Name  interface{}
			ID    string
			Keys  []string
			Fresh bool
		}{name, id, keys, fresh}

		return ctx.Status(fiber.StatusOK).JSON(json)
	}
}

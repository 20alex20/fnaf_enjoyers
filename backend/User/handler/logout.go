package handler

import (
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/gofiber/fiber/v2"
)

// LogOut godoc
// @Summary Logs out user
// @Description Destroys current session.
// @Tags user
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Success 200 {string} string
// @Router /user/logout [post]
func LogOut() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := config.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		if nickname := session.Get("name"); nickname == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		err = session.Destroy()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON("success")
	}
}

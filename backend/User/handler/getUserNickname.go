package handler

import (
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/gofiber/fiber/v2"
)

// GetUserNickname godoc
// @Summary Get current username
// @Description Gets current user's nickname.
// @Tags user
// @Produce json
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Success 200 {object} model.NicknameJSON
// @Router /user/nickname [get]
func GetUserNickname() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := config.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		nickname := session.Get("name")
		if nickname == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		res := model.NicknameJSON{Nickname: nickname.(string)}

		return ctx.Status(fiber.StatusOK).JSON(res)
	}
}

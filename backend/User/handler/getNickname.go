package handler

import (
	"github.com/fnaf-enjoyers/user/config"
	"github.com/fnaf-enjoyers/user/model"
	"github.com/gofiber/fiber/v2"
)

func GetNickname() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := config.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		nickname := session.Get("name")
		if nickname == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		res := model.NicknameResponse{Nickname: nickname.(string)}

		return ctx.Status(fiber.StatusOK).JSON(res)
	}
}

package handler

import (
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/fnaf-enjoyers/user-service/usecase"
	"github.com/gofiber/fiber/v2"
)

func CheckNickname(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		nickname := ctx.Query("nickname", "")
		if nickname == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("nickname is empty")
		}

		thereIs, err := uc.CheckNickname(nickname, repo)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(model.CheckUsernameResponse{ThereIs: thereIs})
	}
}

package handler

import (
	"github.com/fnaf-enjoyers/user/config"
	"github.com/fnaf-enjoyers/user/model"
	"github.com/fnaf-enjoyers/user/repository"
	"github.com/fnaf-enjoyers/user/usecase"
	"github.com/gofiber/fiber/v2"
)

func AuthUser(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req model.RegisterUserRequest
		err := ctx.BodyParser(&req)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		err = uc.AuthUser(req, repo)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
		}

		session, err := config.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		session.Set("name", req.Nickname)
		err = session.Save()
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON("success")
	}
}

package handler

import (
	"github.com/fnaf-enjoyers/user/config"
	"github.com/fnaf-enjoyers/user/model"
	"github.com/fnaf-enjoyers/user/repository"
	"github.com/fnaf-enjoyers/user/usecase"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := config.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		nickname := session.Get("name")
		if nickname == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		var req model.CreatePost
		err = ctx.BodyParser(&req)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		err = uc.CreatePost(nickname.(string), req.Text, req.Categories, req.Filters, repo)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON("success")
	}
}

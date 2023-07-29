package handler

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/fnaf-enjoyers/post-service/pkg/usecase"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req model.CreatePostRequest
		err := ctx.BodyParser(&req)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		nickname := ctx.Query("nickname")
		if nickname == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		err = uc.CreatePost(req.Categories, req.Filters, req.Text, nickname, repo)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON("success")
	}
}

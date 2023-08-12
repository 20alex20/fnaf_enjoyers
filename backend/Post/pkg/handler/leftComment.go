package handler

import (
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

func LeftComment(repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		text := ctx.FormValue("text", "")
		if text == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("text is empty")
		}

		referenceID := ctx.FormValue("reference_id", "")
		if referenceID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("reference_id is empty")
		}

		userID := ctx.FormValue("user_id", "")
		if userID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("user_id is empty")
		}

		err := repo.LeftComment(text, userID, referenceID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		ctx.Status(fiber.StatusOK)
		return nil
	}
}

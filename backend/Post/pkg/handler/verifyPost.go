package handler

import (
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

func VerifyPost(repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		postID := ctx.Query("post_id", "")
		err := repo.VerifyPost(postID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		ctx.Status(fiber.StatusOK)
		return nil
	}
}

package handler

import (
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/fnaf-enjoyers/post-service/pkg/usecase"
	"github.com/gofiber/fiber/v2"
)

func GetComments(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		postID := ctx.Query("post_id", "")
		if postID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("post_id is empty")
		}

		comments := uc.GetCommentsRecursive(postID, repo)

		return ctx.Status(fiber.StatusOK).JSON(comments)
	}
}

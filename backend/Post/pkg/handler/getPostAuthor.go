package handler

import (
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

func GetPostAuthor(repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		postID := ctx.Query("post_id")

		authorID, err := repo.GetPostAuthor(postID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(authorID)
	}
}

package handler

import (
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/fnaf-enjoyers/post-service/pkg/usecases"
	"github.com/gofiber/fiber/v2"
)

// UserPost godoc
// @Summary Return all user posts
// @Param user_id query string true "user id"
// @Router /user-post [get]
func UserPost(ser usecases.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID := ctx.Query("user_id")
		posts, err := ser.UserPost(userID, repo)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}

		return ctx.Status(fiber.StatusOK).JSON(posts)
	}
}

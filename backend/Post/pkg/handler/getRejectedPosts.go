package handler

import (
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/fnaf-enjoyers/post-service/pkg/usecase"
	"github.com/gofiber/fiber/v2"
)

// GetRejectedPosts godoc
// @Summary RejectedPost posts
// @Description Gets current user's rejected posts
// @Tags post
// @Produce json
// @Param id query string true "user id" example("some-uuid")
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Success 200 {array} model.PostUser
// @Router /post/rejected [get]
func GetRejectedPosts(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID := ctx.Query("id")
		if userID == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		res, err := uc.GetRejectedPosts(userID, repo)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(res)
	}
}

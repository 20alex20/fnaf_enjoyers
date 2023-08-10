package handler

import (
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/fnaf-enjoyers/post-service/pkg/usecase"
	"github.com/gofiber/fiber/v2"
)

// GetLikedPosts godoc
// @Summary Liked posts
// @Description Gets current user's liked posts
// @Tags post
// @Produce json
// @Param id query string true "user id" example("some-uuid")
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Success 200 {array} model.PostResponse
// @Router /post/liked [get]
func GetLikedPosts(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID := ctx.Query("id")
		if userID == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		res, err := uc.GetLikedPosts(userID, repo)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(res)
	}
}

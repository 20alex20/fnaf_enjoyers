package handler

import (
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/fnaf-enjoyers/post-service/pkg/usecase"
	"github.com/gofiber/fiber/v2"
)

// GetUserPosts godoc
// @Summary User posts
// @Description Gets current user's posts
// @Tags post
// @Produce json
// @Param nickname query string true "username" example("John Doe")
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Success 200 {array} model.PostUser
// @Router /post/get [get]
func GetUserPosts(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		nickname := ctx.Query("nickname")
		if nickname == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		res, err := uc.GetUserPosts(nickname, repo)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(res)
	}
}

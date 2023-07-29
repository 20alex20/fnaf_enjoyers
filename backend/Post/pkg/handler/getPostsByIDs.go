package handler

import (
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/fnaf-enjoyers/post-service/pkg/usecase"
	"github.com/gofiber/fiber/v2"
)

// GetPostsByIDs godoc
// @Summary Get posts by IDs
// @Description Gets posts from current array of IDs
// @Tags post
// @Accept json
// @Produce json
// @Param req body []string true "request body"
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Success 200 {array} model.PostUser
// @Router /post/create [post]
func GetPostsByIDs(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var IDs []string
		err := ctx.BodyParser(&IDs)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		posts, err := uc.GetPostsByIDs(IDs, repo)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(posts)
	}
}

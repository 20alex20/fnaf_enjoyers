package handler

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

// CreatePost godoc
// @Summary Create post
// @Description Creates post from current user
// @Tags post
// @Accept json
// @Produce json
// @Param nickname query string true "username" example("John Doe")
// @Param req body model.CreatePostRequest true "request body"
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Failure 400 {string} string
// @Success 200 {array} model.PostResponse
// @Router /post/create [post]
func CreatePost(repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req model.CreatePostRequest
		err := ctx.BodyParser(&req)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		userID := ctx.Query("id")
		if userID == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		err = repo.CreatePost(req.Categories, req.Filters, req.Text, userID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON("success")
	}
}

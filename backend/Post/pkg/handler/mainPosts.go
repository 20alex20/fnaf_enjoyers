package handler

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/fnaf-enjoyers/post-service/pkg/usecases"
	"github.com/gofiber/fiber/v2"
)

// MainPosts godoc
// @Summary Main page posts
// @Description Gets posts on the main page
// @Tags main
// @Accept json
// @Produce json
// @Param requestBody body model.MainPostRequest true "request body"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Success 200 {array} model.PostMain
// @Router /main/posts [get]
func MainPosts(uc usecases.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var requestBody model.MainPostRequest
		err := ctx.BodyParser(&requestBody)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		posts, err := uc.MainPosts(requestBody, repo)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(posts)
	}
}

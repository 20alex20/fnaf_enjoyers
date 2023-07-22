package handler

import (
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
// @Failure 500 {string} string
// @Success 200 {array} model.PostMain
// @Router /main/posts [get]
func MainPosts(uc usecases.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		category := ctx.Query("category", "all")
		filter := ctx.Query("filter", "without")
		sort := ctx.Query("sort", "id")
		number := ctx.QueryInt("number", 2)
		page := ctx.QueryInt("page", 1)

		posts, err := uc.MainPosts(category, filter, sort, number, page, repo)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(posts)
	}
}

package handler

import (
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/fnaf-enjoyers/post-service/pkg/usecase"
	"github.com/gofiber/fiber/v2"
)

// MainPosts godoc
// @Summary Main page posts
// @Description Gets posts on the main page
// @Tags main
// @Produce json
// @Param category query string true "post category" example(all)
// @Param filter query string true "post filter" example(without)
// @Param sort query string true "sort type" example(date_time)
// @Param number query int true "total number of posts to get" example(2)
// @Param page query int true "page number" "1" example(1)
// @Failure 500 {string} string
// @Success 200 {array} model.PostMain
// @Router /main/posts [get]
func MainPosts(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
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

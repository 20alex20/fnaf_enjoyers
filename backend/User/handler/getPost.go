package handler

import (
	"encoding/json"
	"fmt"
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/fnaf-enjoyers/user-service/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// GetPost godoc
// @Summary Get post
// @Description Gets post by id
// @Tags post
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Success 200 {object} model.PostUser
// @Router /post [get]
func GetPost(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		postID := ctx.Query("post_id", "")
		if postID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("post_id is empty")
		}

		url := fmt.Sprintf("http://localhost:3001/post?post_id=%s", postID)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		var post model.PostResponse
		if err = json.NewDecoder(res.Body).Decode(&post); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		url = fmt.Sprintf("http://localhost:3001/post/incr-view?post_id=%s", postID)
		req, err = http.NewRequest("POST", url, nil)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		_, err = http.DefaultClient.Do(req)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		authorID, err := uc.GetPostAuthor(postID)

		err = repo.IncrViews(authorID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(post)
	}
}

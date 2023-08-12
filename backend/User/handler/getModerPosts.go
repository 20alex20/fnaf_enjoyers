package handler

import (
	"encoding/json"
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// GetModerPosts godoc
// @Summary Get moder posts
// @Description Gets posts from current user
// @Tags post
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Failure 403 {string} string
// @Success 200 {array} model.PostUser
// @Router /posts/moder [get]
func GetModerPosts(repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := config.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		nickname := session.Get("name")
		if nickname == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		isModerator, err := repo.CheckIfModerator(nickname.(string))
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		if !isModerator {
			return ctx.Status(fiber.StatusForbidden).JSON("operation cannot be performed by a non-moderator")
		}

		postReqUrl := "http://localhost:3001/posts/moder"
		postReq, err := http.NewRequest("GET", postReqUrl, nil)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		res, err := http.DefaultClient.Do(postReq)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		posts := make([]model.ModerPostResponse, 0)
		if err = json.NewDecoder(res.Body).Decode(&posts); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(posts)
	}
}

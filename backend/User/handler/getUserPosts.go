package handler

import (
	"encoding/json"
	"fmt"
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// GetUserPosts godoc
// @Summary Get user posts
// @Description Gets posts from current user
// @Tags post
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Success 200 {array} model.PostUser
// @Router /posts [get]
func GetUserPosts(repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := config.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		nickname := session.Get("name")
		if nickname == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		userID, err := repo.GetUserID(nickname.(string))
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		postReqUrl := fmt.Sprintf("http://localhost:3001/posts?id=%s", userID)
		postReq, err := http.NewRequest("GET", postReqUrl, nil)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		res, err := http.DefaultClient.Do(postReq)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		posts := make([]model.PostResponse, 0)
		if err = json.NewDecoder(res.Body).Decode(&posts); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(posts)
	}
}

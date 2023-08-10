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

func GetLikedPosts(repo repository.Repository) fiber.Handler {
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

		postReqUrl := fmt.Sprintf("http://localhost:3001/post/liked?id=%s", userID)
		postReq, err := http.NewRequest("GET", postReqUrl, nil)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		res, err := http.DefaultClient.Do(postReq)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		var posts []model.PostResponse
		if err = json.NewDecoder(res.Body).Decode(&posts); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(posts)
	}
}

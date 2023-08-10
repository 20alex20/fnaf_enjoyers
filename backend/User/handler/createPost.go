package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/fnaf-enjoyers/user-service/model"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// CreatePost godoc
// @Summary Create post
// @Description Creates post from current user
// @Tags post
// @Accept json
// @Produce json
// @Param req body model.CreatePostRequest true "request body"
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Failure 400 {string} string
// @Success 200 {array} model.PostResponse
// @Router /post/create [post]
func CreatePost(repo repository.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := config.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		nickname := session.Get("name")
		if nickname == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		var req model.CreatePostRequest
		err = ctx.BodyParser(&req)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		userID, err := repo.GetUserID(nickname.(string))
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		postReqUrl := fmt.Sprintf("http://localhost:3001/post/create?id=%s", userID)
		var buf bytes.Buffer
		err = json.NewEncoder(&buf).Encode(req)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		postReq, err := http.NewRequest("POST", postReqUrl, &buf)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		postReq.Header.Set("Content-Type", "application/json")

		res, err := http.DefaultClient.Do(postReq)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		if res.Status != "200 OK" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(res.Body)
		}

		err = repo.IncrPosts(nickname.(string))
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON("success")
	}
}

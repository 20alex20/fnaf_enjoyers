package handler

import (
	"fmt"
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// RejectPost godoc
// @Summary Reject posts
// @Description Reject post by moderator
// @Tags post
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Failure 403 {string} string
// @Failure 400 {string} string
// @Success 200 {array} model.PostUser
// @Router /post/verify [post]
func RejectPost(repo repository.Repository) fiber.Handler {
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

		postID := ctx.FormValue("post_id", "")
		if postID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("post_id is empty")
		}

		text := ctx.FormValue("text", "")
		if text == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON("text is empty")
		}

		postReqUrl := fmt.Sprintf("http://localhost:3001/post/reject?post_id=%s&text=%s", postID, text)
		postReq, err := http.NewRequest("POST", postReqUrl, nil)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		_, err = http.DefaultClient.Do(postReq)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		ctx.Status(fiber.StatusOK)
		return nil
	}
}

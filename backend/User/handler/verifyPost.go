package handler

import (
	"fmt"
	"github.com/fnaf-enjoyers/user-service/config"
	"github.com/fnaf-enjoyers/user-service/repository"
	"github.com/fnaf-enjoyers/user-service/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// VerifyPost godoc
// @Summary Verify posts
// @Description Verify post by moderator
// @Tags post
// @Failure 500 {string} string
// @Failure 401 {string} string
// @Failure 403 {string} string
// @Failure 400 {string} string
// @Success 200 {array} model.PostUser
// @Router /post/verify [post]
func VerifyPost(uc usecase.UseCase, repo repository.Repository) fiber.Handler {
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

		postReqUrl := fmt.Sprintf("http://localhost:3001/post/verify?post_id=%s", postID)
		postReq, err := http.NewRequest("POST", postReqUrl, nil)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		_, err = http.DefaultClient.Do(postReq)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		authorID, err := uc.GetPostAuthor(postID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		err = repo.IncrPosts(authorID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}

		ctx.Status(fiber.StatusOK)
		return nil
	}
}

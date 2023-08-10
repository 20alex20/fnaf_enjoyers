package repository

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/jackc/pgx"
)

func (r *repository) GetPostByUserID(userID string) ([]model.PostDTO, error) {
	var posts []model.PostDTO

	err := r.DB.Select(&posts, GetUserPosts, userID)
	if err != nil {
		log.Errorf("Unable to receive user posts: %s", err)
		return nil, err
	}

	return posts, nil
}

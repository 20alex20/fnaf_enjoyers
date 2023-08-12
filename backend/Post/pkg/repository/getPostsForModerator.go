package repository

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/gofiber/fiber/v2/log"
)

func (r *repository) GetPostsForModerator() ([]model.PostDTO, error) {
	var posts []model.PostDTO

	err := r.DB.Select(&posts, GetModerPosts)
	if err != nil {
		log.Errorf("Unable to receive moder posts: %s", err)
		return nil, err
	}

	return posts, nil
}

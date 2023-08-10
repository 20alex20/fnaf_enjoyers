package repository

import "github.com/fnaf-enjoyers/post-service/pkg/model"

func (r *repository) GetRejectedPosts(userID string) ([]model.RejectedPostDTO, error) {
	var posts []model.RejectedPostDTO

	err := r.DB.Select(&posts, RejectedPost, userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

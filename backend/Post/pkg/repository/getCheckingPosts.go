package repository

import "github.com/fnaf-enjoyers/post-service/pkg/model"

func (r *repository) GetCheckingPosts(userID string) ([]model.CheckingPostDTO, error) {
	var posts []model.CheckingPostDTO

	err := r.DB.Select(&posts, CheckingPost, userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

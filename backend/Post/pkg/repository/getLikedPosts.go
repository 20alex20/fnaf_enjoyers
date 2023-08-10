package repository

import "github.com/fnaf-enjoyers/post-service/pkg/model"

func (r *repository) GetLikedPosts(userID string) ([]model.PostDTO, error) {
	var posts []model.PostDTO

	err := r.DB.Select(&posts, Liked, userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

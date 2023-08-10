package repository

import "github.com/fnaf-enjoyers/post-service/pkg/model"

func (r *repository) GetPost(postID string) (model.PostDTO, error) {
	var post model.PostDTO
	err := r.DB.Get(&post, Post, postID)
	if err != nil {
		return model.PostDTO{}, err
	}

	return post, nil
}

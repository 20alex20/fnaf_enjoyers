package repository

import "github.com/fnaf-enjoyers/post-service/pkg/model"

func (r *repository) GetComments(referenceID string) ([]model.CommentDTO, error) {
	res := make([]model.CommentDTO, 0)
	err := r.DB.Select(&res, GetComments, referenceID)
	if err != nil {
		return res, err
	}

	return res, nil
}

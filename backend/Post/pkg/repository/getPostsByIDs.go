package repository

import "github.com/fnaf-enjoyers/post-service/pkg/model"

func (r *repository) GetPostsByIDs(IDs []string) ([]model.PostDTO, error) {
	var postsDTO []model.PostDTO
	query := `
		select * from post where id in $1;
`
	err := r.DB.Select(&postsDTO, query, IDs)
	if err != nil {
		return nil, err
	}

	return postsDTO, nil
}

package usecase

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

func (uc *useCase) GetPostsByIDs(IDs []string, repo repository.Repository) ([]model.PostUser, error) {
	postsDTO, err := repo.GetPostsByIDs(IDs)
	if err != nil {
		return nil, err
	}

	var posts []model.PostUser

	for id, postDTO := range postsDTO {
		post := model.PostUser{
			ID:       id,
			Text:     postDTO.Text,
			DateTime: postDTO.Date.Format("2006.01.02 15:04"),
			Views:    postDTO.Views,
			Likes:    postDTO.Likes,
		}

		posts = append(posts, post)
	}

	return posts, nil
}

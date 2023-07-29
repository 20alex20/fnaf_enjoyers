package usecase

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

func (uc *useCase) GetUserPosts(nickname string, repo repository.Repository) ([]model.PostUser, error) {
	postsDTO, err := repo.GetPostByNickname(nickname)
	if err != nil {
		return nil, err
	}

	var res []model.PostUser

	for i, postDTO := range postsDTO {
		post := model.PostUser{
			ID:       i + 1,
			Text:     postDTO.Text,
			DateTime: postDTO.Date.Format("2006.01.02 15:04"),
			Views:    postDTO.Views,
			Likes:    postDTO.Likes,
		}

		res = append(res, post)
	}

	return res, nil
}

package usecase

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

func (uc *useCase) GetPost(postID string, repo repository.Repository) (model.PostResponse, error) {
	postDTO, err := repo.GetPost(postID)
	if err != nil {
		return model.PostResponse{}, err
	}

	post := model.PostResponse{
		ID:       postDTO.ID,
		Text:     postDTO.Text,
		DateTime: postDTO.Date.Format("2006.01.02 15:04"),
		Views:    postDTO.Views,
		Likes:    postDTO.Likes,
	}

	return post, nil
}

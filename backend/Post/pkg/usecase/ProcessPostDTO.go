package usecase

import "github.com/fnaf-enjoyers/post-service/pkg/model"

func (uc *useCase) ProcessPostDTO(postsDTO []model.PostDTO) (posts []model.PostResponse) {
	for _, postDTO := range postsDTO {
		post := model.PostResponse{
			ID:       postDTO.ID,
			Text:     postDTO.Text,
			DateTime: postDTO.Date.Format("2006.01.02 15:04"),
			Views:    postDTO.Views,
			Likes:    postDTO.Likes,
		}

		posts = append(posts, post)
	}

	return posts
}

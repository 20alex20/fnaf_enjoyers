package usecases

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

func (s *service) UserPost(userID string, repo repository.Repository) ([]model.Post, error) {
	postsDTO, err := repo.GetPostByUserID(userID)
	if err != nil {
		return nil, err
	}

	//req, err := http.NewRequest("GET", "http://localhost:3002/username?")

	var posts []model.Post

	for _, postDTO := range postsDTO {
		post := model.Post{
			Username: "Placeholder",
			Date:     postDTO.Date.Format("2006.01.02 15:04"),
			Text:     postDTO.Text,
			Views:    postDTO.Views,
			Likes:    postDTO.Likes,
		}

		posts = append(posts, post)
	}

	return posts, nil
}

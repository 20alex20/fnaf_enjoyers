package usecases

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

// MainPosts fills main page with posts
// with certain criteria from db
//
// - requestBody: json containing search parameters
func (s *service) MainPosts(requestBody model.MainPostRequest, repo repository.Repository) ([]model.PostMain, error) {
	filtered := requestBody.Filter != "without"
	categorized := requestBody.Category != "all"

	order := "id"

	switch requestBody.Sort {
	case "date_time":
		order = model.DateTime
		break
	case "likes":
		order = model.Likes
		break
	case "views":
		order = model.Views
	}

	var postsDTO []model.PostDTO
	var err error
	err = nil

	if filtered && categorized {
		postsDTO, err = repo.GetMainPostsCF(
			requestBody.Category,
			requestBody.Filter,
			order,
			requestBody.Number,
			requestBody.Page)
	} else if filtered {
		postsDTO, err = repo.GetMainPostsFiltered(
			requestBody.Filter,
			order,
			requestBody.Number,
			requestBody.Page)
	} else if categorized {
		postsDTO, err = repo.GetMainPostsCategorized(
			requestBody.Category,
			order,
			requestBody.Number,
			requestBody.Page)
	} else {
		postsDTO, err = repo.GetMainPosts(
			order,
			requestBody.Number,
			requestBody.Page)
	}

	if err != nil {
		return nil, err
	}

	var posts []model.PostMain

	for _, postDTO := range postsDTO {
		post := model.PostMain{
			Text:     postDTO.Text,
			DateTime: postDTO.Date.Format("2006.01.02 15:04"),
			Views:    postDTO.Views,
			Likes:    postDTO.Likes,
		}

		posts = append(posts, post)
	}

	return posts, nil
}

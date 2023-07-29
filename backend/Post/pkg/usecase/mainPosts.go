package usecase

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

// MainPosts fills main page with posts
// with certain criteria from db
//
// - category: category to search from
//
// - filter: faculty to search from
//
// - sort: type of post sorting (date_time, likes, views)
//
// - number: total number of posts (defining max page)
//
// - page: defines an offset for sql scrolling
func (uc *useCase) MainPosts(category, filter, sort string, number, page int, repo repository.Repository) ([]model.PostMain, error) {
	filtered := filter != "without"
	categorized := category != "all"

	order := "id"

	switch sort {
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
			category,
			filter,
			order,
			number,
			page)
	} else if filtered {
		postsDTO, err = repo.GetMainPostsFiltered(
			filter,
			order,
			number,
			page)
	} else if categorized {
		postsDTO, err = repo.GetMainPostsCategorized(
			category,
			order,
			number,
			page)
	} else {
		postsDTO, err = repo.GetMainPosts(
			order,
			number,
			page)
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

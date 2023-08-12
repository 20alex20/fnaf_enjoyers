package repository

import (
	"fmt"
	"github.com/fnaf-enjoyers/post-service/pkg/model"
)

// GetMainPosts gets top number posts in pages of 2 posts
//
// - order: type of post sorting (date, likes, views)
//
// - number: total number of posts (defining max page)
//
// - page: defines an offset for sql scrolling
func (r *repository) GetMainPosts(order string, number, page int) ([]model.PostDTO, error) {
	var posts []model.PostDTO

	offset := number * (page - 1)
	query := fmt.Sprintf("select * from post where checked = true and accepted = true order by %s desc limit $1 offset $2;", order)

	err := r.DB.Select(&posts, query, number, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

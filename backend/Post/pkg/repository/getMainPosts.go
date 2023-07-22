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

	r.PageCorrection(&page, number, 2)

	offset := 2 * (page - 1)
	query := fmt.Sprintf("select * from post order by %s desc limit 2 offset $1", order)

	err := r.DB.Select(&posts, query, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

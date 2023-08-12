package repository

import (
	"fmt"
	"github.com/fnaf-enjoyers/post-service/pkg/model"
)

// GetMainPostsFiltered gets top number posts in pages of 2 posts
// with filter support
//
// - filter: faculty to search from
//
// - order: type of post sorting (date, likes, views)
//
// - number: total number of posts (defining max page)
//
// - page: defines an offset for sql scrolling
func (r *repository) GetMainPostsFiltered(filter, order string, number, page int) ([]model.PostDTO, error) {
	var posts []model.PostDTO

	offset := number * (page - 1)
	query := fmt.Sprintf(`
		select * from post 
			where id in 
				(select post_id from post_filter where filter_id in (
		                   select id from filter where title = $1
		               )) 
			and checked = true
			and accepted = true
		order by %s desc 
		limit $2 offset $3;
	`, order)

	err := r.DB.Select(&posts, query, filter, number, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

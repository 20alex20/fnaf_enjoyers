package repository

import (
	"fmt"
	"github.com/fnaf-enjoyers/post-service/pkg/model"
)

// GetMainPostsCategorized gets top number posts in pages of 2 posts
// with category support
//
// - category: category to search from
//
// - order: type of post sorting (date_time, likes, views)
//
// - number: total number of posts (defining max page)
//
// - page: defines an offset for sql scrolling
func (r *repository) GetMainPostsCategorized(category, order string, number, page int) ([]model.PostDTO, error) {
	categoryIDs, err := r.GetIDsByTitles("category", []string{category})
	if err != nil {
		return nil, err
	}

	var posts []model.PostDTO

	r.PageCorrection(&page, number, 2)

	offset := 2 * (page - 1)
	query := fmt.Sprintf(`
		select * from post 
		         where id in 
		               (select post_id from post_category where category_id in $1) 
		                   order by %s desc 
		                   limit 2 offset $2
	`, order)

	err = r.DB.Select(&posts, query, categoryIDs, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

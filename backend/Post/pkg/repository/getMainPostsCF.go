package repository

import (
	"fmt"
	"github.com/fnaf-enjoyers/post-service/pkg/model"
)

// GetMainPostsCF gets top number posts in pages of 2 posts
// with both category and filter support
//
// - category: category to search from
//
// - filter: faculty to search from
//
// - order: type of post sorting (date, likes, views)
//
// - number: total number of posts (defining max page)
//
// - page: defines an offset for sql scrolling
func (r *repository) GetMainPostsCF(category, filter, order string, number, page int) ([]model.PostDTO, error) {
	var postIDs []string

	query := `
		select pf.post_id 
		from post_filter pf 
		    inner join post_category pc 
		        on pc.post_id = pf.post_id 
		where 
		    pf.filter_id in
			(select id from filter where title = $1)
		and 
		    pc.category_id in
			(select id from category where title = $2)
	`
	err := r.DB.Select(&postIDs, query, filter, category)
	if err != nil {
		return nil, err
	}

	var posts []model.PostDTO

	r.PageCorrection(&page, number, 2)

	offset := 2 * (page - 1)
	query = fmt.Sprintf("select * from post where id in $1 order by %s desc limit 2 offset $2", order)

	err = r.DB.Select(&posts, query, postIDs, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

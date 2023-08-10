package repository

func (r *repository) GetPostFilters(postID string) ([]string, error) {
	filters := make([]string, 0)
	err := r.DB.Select(&filters, PostFilters, postID)
	if err != nil {
		return nil, err
	}
	return filters, nil
}

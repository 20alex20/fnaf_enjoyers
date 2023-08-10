package repository

func (r *repository) GetPostCategories(postID string) ([]string, error) {
	categories := make([]string, 0)
	err := r.DB.Select(&categories, PostCategories, postID)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

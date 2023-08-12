package repository

func (r *repository) CreatePost(categories, filters []string, text, userID string) error {
	var postID string
	err := r.DB.Get(&postID, CreatePostReturnID, userID, text)
	if err != nil {
		return err
	}

	for _, category := range categories {
		_, err = r.DB.Exec(CreateCategoryRelation, postID, category)
		if err != nil {
			return err
		}
	}

	for _, filter := range filters {
		_, err = r.DB.Exec(CreateFilterRelation, postID, filter)
		if err != nil {
			return err
		}
	}

	return nil
}

package repository

func (r *repository) CheckLike(userID, postID string) (bool, error) {
	var liked bool
	query := `
		select exists(select 1 from user_post_liked where user_id = $1 and post_id = $2);
`
	err := r.DB.Get(&liked, query, userID, postID)
	if err != nil {
		return false, err
	}

	return liked, nil
}

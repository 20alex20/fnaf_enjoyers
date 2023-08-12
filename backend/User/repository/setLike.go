package repository

func (r *repository) SetLike(userID, postID string) error {
	query := `
		insert into user_post_liked
		(user_id, post_id)
		values ($1, $2);
`
	_, err := r.DB.Exec(query, userID, postID)
	return err
}

package repository

func (r *repository) UnsetLike(userID, postID string) error {
	query := `
		delete from user_post_liked
		where user_id=$1 and post_id=$2;
`
	_, err := r.DB.Exec(query, userID, postID)
	return err
}

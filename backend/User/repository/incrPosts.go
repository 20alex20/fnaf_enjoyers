package repository

func (r *repository) IncrPosts(userID string) error {
	query := `
		update "user"
		set user_posts = user_posts + 1
		where id=$1;
`
	_, err := r.DB.Exec(query, userID)
	return err
}

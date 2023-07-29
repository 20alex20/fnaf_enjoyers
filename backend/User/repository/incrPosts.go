package repository

func (r *repository) IncrPosts(nickname string) error {
	query := `
		update "user"
		set user_posts = user_posts + 1
		where nickname=$1;
`
	_, err := r.DB.Exec(query, nickname)
	return err
}

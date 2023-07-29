package repository

func (r *repository) GetLikedPostsIDs(nickname string) ([]string, error) {
	IDs := make([]string, 0)

	query := `
		select post_id 
		from user_post_liked 
		where user_id = 
		      (select id from "user" where nickname = $1);
`

	err := r.DB.Select(&IDs, query, nickname)
	if err != nil {
		return IDs, err
	}

	return IDs, nil
}

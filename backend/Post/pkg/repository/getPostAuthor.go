package repository

func (r *repository) GetPostAuthor(postID string) (string, error) {
	var authorID string
	err := r.DB.Get(&authorID, GetPostAuthor, postID)
	if err != nil {
		return "", err
	}

	return authorID, nil
}

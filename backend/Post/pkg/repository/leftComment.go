package repository

func (r *repository) LeftComment(text, userID, referenceID string) error {
	_, err := r.DB.Exec(LeftComment, referenceID, userID, text)
	return err
}

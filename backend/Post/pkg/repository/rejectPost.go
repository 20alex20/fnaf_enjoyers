package repository

func (r *repository) RejectPost(postID, text string) error {
	_, err := r.DB.Exec(RejectPost, postID)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(InsertRejectedMessage, postID, text)
	return err
}

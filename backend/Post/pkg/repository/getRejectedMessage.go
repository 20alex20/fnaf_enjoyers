package repository

func (r *repository) GetRejectedMessage(postID string) (string, error) {
	message := ""
	err := r.DB.Get(&message, RejectedMessage, postID)
	if err != nil {
		return "", err
	}
	return message, nil
}

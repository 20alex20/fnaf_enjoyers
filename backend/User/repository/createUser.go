package repository

func (r *repository) CreateUser(nickname, password string) error {
	query := `
		insert into "user"
			(nickname, password)
		values (
		        $1, $2
		);
	`
	_, err := r.DB.Exec(query, nickname, password)
	if err != nil {
		return err
	}

	return nil
}

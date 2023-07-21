package repository

// GetIDsByTitles gets an array of IDs from any table
// that contains title and id fields
//
// - table: table name to search in
//
// - titles: array of titles to get IDs for
func (r *repository) GetIDsByTitles(table string, titles []string) ([]string, error) {
	var ids []string
	err := r.DB.Select(&ids, "select id from $1 where title in $2", table, titles)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

package repository

func (r *repository) GetMaxPage(number int) (int, error) {
	var total int
	err := r.DB.Get(&total, CountMainPosts)
	if err != nil {
		return 0, err
	}

	div := total / number
	if total%number != 0 {
		div++
	}

	return div, nil
}

func (r *repository) GetMaxPageCategorized(number int, category string) (int, error) {
	var total int
	err := r.DB.Get(&total, CountMainPostsCategorized, category)
	if err != nil {
		return 0, err
	}

	div := total / number
	if total%number != 0 {
		div++
	}

	return div, nil
}

func (r *repository) GetMaxPageFiltered(number int, filter string) (int, error) {
	var total int
	err := r.DB.Get(&total, CountMainPostsFiltered, filter)
	if err != nil {
		return 0, err
	}

	div := total / number
	if total%number != 0 {
		div++
	}

	return div, nil
}

func (r *repository) GetMaxPageCF(number int, filter, category string) (int, error) {
	var total int
	err := r.DB.Get(&total, CountMainPostsCF, filter, category)
	if err != nil {
		return 0, err
	}

	div := total / number
	if total%number != 0 {
		div++
	}

	return div, nil
}

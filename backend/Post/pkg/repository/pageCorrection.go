package repository

func (r *repository) PageCorrection(page *int, number, limit int) {
	additionalPage := 0
	if dividable := number%limit == 0; !dividable {
		additionalPage++
	}

	if *page < 1 {
		*page = 1
	} else if *page > (number/limit + additionalPage) {
		*page = number/limit + additionalPage
	}
}

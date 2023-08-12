package usecase

import (
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

func (uc *useCase) MaxPage(category, filter string, number int, repo repository.Repository) (int, error) {
	filtered := filter != "without"
	categorized := category != "all"

	maxPage := 0
	var err error
	err = nil

	if filtered && categorized {
		maxPage, err = repo.GetMaxPageCF(number, filter, category)
	} else if filtered {
		maxPage, err = repo.GetMaxPageFiltered(number, filter)
	} else if categorized {
		maxPage, err = repo.GetMaxPageCategorized(number, category)
	} else {
		maxPage, err = repo.GetMaxPage(number)
	}

	return maxPage, err
}

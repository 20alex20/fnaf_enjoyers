package usecase

import "github.com/fnaf-enjoyers/user-service/repository"

func (uc *useCase) CheckNickname(nickname string, repo repository.Repository) (bool, error) {
	thereIs, err := repo.CheckUser(nickname)
	if err != nil {
		return false, err
	}

	return thereIs, nil
}

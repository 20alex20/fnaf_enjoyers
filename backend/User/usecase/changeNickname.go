package usecase

import "github.com/fnaf-enjoyers/user-service/repository"

func (uc *useCase) ChangeNickname(old, new string, repo repository.Repository) error {
	err := repo.ChangeNickname(old, new)
	return err
}

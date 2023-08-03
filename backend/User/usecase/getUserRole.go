package usecase

import "github.com/fnaf-enjoyers/user-service/repository"

func (uc *useCase) GetUserRole(nickname string, repo repository.Repository) (string, error) {
	isModerator, err := repo.GetUserRole(nickname)

	if err != nil {
		return "", err
	}

	if *isModerator {
		return "admin", nil
	}

	return "user", nil
}

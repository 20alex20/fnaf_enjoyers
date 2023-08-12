package usecase

import (
	"fmt"
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
)

func (uc *useCase) GetCommentsRecursive(referenceID string, repo repository.Repository) []model.CommentResponse {
	commentsDTO, err := repo.GetComments(referenceID)
	if err != nil {
		fmt.Println(err.Error())
	}

	var res []model.CommentResponse

	for _, commentDTO := range commentsDTO {
		author, err := uc.GetNickname(commentDTO.UserID)
		if err != nil {
			fmt.Println(err.Error())
		}

		profilePic, err := uc.GetProfilePic(commentDTO.UserID)
		if err != nil {
			fmt.Println(err.Error())
		}

		replies := uc.GetCommentsRecursive(commentDTO.ID, repo)

		comment := model.CommentResponse{
			ID:      commentDTO.ID,
			Text:    commentDTO.Text,
			Date:    commentDTO.Date.Format("2006.01.02 15:04"),
			Author:  author,
			Img:     profilePic,
			Replies: replies,
		}

		res = append(res, comment)
	}

	return res
}

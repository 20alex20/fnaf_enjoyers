package repository

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/jackc/pgx"
)

func (r *repository) GetPostByNickname(nickname string) ([]model.PostDTO, error) {
	var posts []model.PostDTO

	query := `
		select * from post where user_id = (select id from "user" where nickname=$1)
		order by date desc;
`

	err := r.DB.Select(&posts, query, nickname)
	if err != nil {
		log.Errorf("Unable to receive user posts: %s", err)
		return nil, err
	}

	return posts, nil
}

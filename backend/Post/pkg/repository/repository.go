package repository

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetPostByNickname(userID string) ([]model.PostDTO, error)
	GetMainPosts(order string, number, page int) ([]model.PostDTO, error)
	GetMainPostsFiltered(filter, order string, number, page int) ([]model.PostDTO, error)
	GetMainPostsCategorized(category, order string, number, page int) ([]model.PostDTO, error)
	GetMainPostsCF(category, filter, order string, number, page int) ([]model.PostDTO, error)
	CreatePost(categories, filters []string, text, nickname string) error
}

type repository struct {
	DB sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		DB: *db,
	}
}

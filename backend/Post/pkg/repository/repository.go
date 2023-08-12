package repository

import (
	"github.com/fnaf-enjoyers/post-service/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetPostByUserID(userID string) ([]model.PostDTO, error)
	GetMainPosts(order string, number, page int) ([]model.PostDTO, error)
	GetMainPostsFiltered(filter, order string, number, page int) ([]model.PostDTO, error)
	GetMainPostsCategorized(category, order string, number, page int) ([]model.PostDTO, error)
	GetMainPostsCF(category, filter, order string, number, page int) ([]model.PostDTO, error)
	CreatePost(categories, filters []string, text, nickname string) error
	GetRejectedPosts(userID string) ([]model.RejectedPostDTO, error)
	GetLikedPosts(userID string) ([]model.PostDTO, error)
	GetPostCategories(postID string) ([]string, error)
	GetPostFilters(postID string) ([]string, error)
	GetRejectedMessage(postID string) (string, error)
	GetPostsForModerator() ([]model.PostDTO, error)
	GetPost(postID string) (model.PostDTO, error)
	VerifyPost(postID string) error
	RejectPost(postID, text string) error
	SetLike(postID string) error
	IncrView(postID string) error
	GetPostAuthor(postID string) (string, error)
	UnsetLike(postID string) error
	GetComments(referenceID string) ([]model.CommentDTO, error)
	LeftComment(text, userID, referenceID string) error
	GetMaxPage(number int) (int, error)
	GetMaxPageCategorized(number int, category string) (int, error)
	GetMaxPageFiltered(number int, filter string) (int, error)
	GetMaxPageCF(number int, filter, category string) (int, error)
}

type repository struct {
	DB sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		DB: *db,
	}
}

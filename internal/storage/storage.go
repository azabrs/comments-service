package storage

import (
	"comments_service/graph/model"
	"comments_service/internal/models"
)

type Storage interface{
	Register(string) error
	CreatePost(models.Post) error
	IsLoginExist(string) error
	Posts(int) ([]*model.Post, error)
	AddComment(model.SComment, []chan *model.RComment, []int) error
	PostAndComment(postID *string, limit *int) (*model.PostWithComment, error)
	IsCommentable(postID *string)(bool, error)
}
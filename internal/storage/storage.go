package storage

import (
	"comments_service/graph/model"
	"comments_service/internal/models"
)
//go:generate go run github.com/vektra/mockery/v2@v2.43.2 --name=Storage
type Storage interface{
	Register(string) error
	CreatePost(models.Post) error
	IsLoginExist(string) error
	Posts(int, int) ([]*model.Post, error)
	AddComment(model.SComment, []chan *model.RComment, []int) error
	PostAndComment(*string, int, int) (*model.PostWithComment, error)
	IsCommentable(*string)(bool, error)
}
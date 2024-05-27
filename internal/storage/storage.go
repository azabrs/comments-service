package storage

import (
	"comments_service/graph/model"
	"comments_service/internal/models"
)

type Storage interface{
	Register(string) error
	CreatePost(models.Post) error
	IsLoginExist(string) error
	Posts() ([]*model.Post, error)
}
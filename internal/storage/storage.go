package storage

import "comments_service/internal/models"

type Storage interface{
	Register(string) error
	CreatePost(models.Post) error
	IsLoginExist(string) error
}
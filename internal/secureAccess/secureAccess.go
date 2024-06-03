package secure_access

import (
	"comments_service/graph/model"
	"comments_service/internal/storage"
)

//go:generate go run github.com/vektra/mockery/v2@v2.43.2 --name=SecureAccess
type SecureAccess interface{
	Authorize(login string) (string, error)
	Authentication(IdentificationData model.IdentificationData, stor storage.Storage) error
}
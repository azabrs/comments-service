package commentusecase

import (
	"comments_service/graph/model"
	"time"

	"comments_service/internal/models"
	secure_access "comments_service/internal/secureAccess"
	"comments_service/internal/storage"
	"context"
)

type UseCase struct{
	stor storage.Storage
	secure secure_access.SecureAccess
}
func New(stor storage.Storage, secure secure_access.SecureAccess) UseCase{
	return UseCase{stor : stor,
						 secure : secure}
}

func (u UseCase)Register(ctx context.Context, registerData model.RegisterData) (string, error){
	token, err := u.secure.Authorize(registerData.Login)
	if err != nil{
		return "", err
	}

	err = u.stor.Register(registerData.Login)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u UseCase) CreatePost(ctx context.Context, identificationData model.IdentificationData, postData string, isCommentEnbale *bool) error{
	if err := u.secure.Authentication(identificationData.Login, identificationData.Token, u.stor); err != nil{
		return err
	}
	Post := models.Post{
		Author : identificationData.Login,
		TimeAdd: time.Now(),
		Subject: postData,
		IsCommentEnable: *isCommentEnbale,
	}
	if err := u.stor.CreatePost(Post); err != nil{
		return err
	}
	return nil
}

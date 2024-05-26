package commentusecase

import (
	"comments_service/graph/model"

	"comments_service/internal/authorization"
	"comments_service/internal/storage"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type UseCase struct{
	stor storage.Storage
	auth authorization.Authorization
}
func New(stor storage.Storage, auth authorization.Authorization) UseCase{
	return UseCase{stor : stor,
						 auth : auth}
}

func (u UseCase)Register(ctx context.Context, registerData model.RegisterData) (string, error){
	token, err := u.auth.Authorize()
	if err != nil{
		return "", err
	}
	tokenHash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	err = u.stor.Register(registerData.Login, string(tokenHash))
	if err != nil {
		return "", err
	}
	return token, nil
}

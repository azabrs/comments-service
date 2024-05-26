package commentusercase

import (
	"comments_service/graph/model"

	"comments_service/internal/authorization"
	"comments_service/internal/storage"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type UserCase struct{
	stor storage.Storage
	auth authorization.Authorization
}

func (u UserCase)Register(ctx context.Context, registerData model.RegisterData) (string, error){
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

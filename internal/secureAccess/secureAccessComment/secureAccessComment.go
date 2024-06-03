package secure_access_comment

import (
	custom_errors "comments_service/errors"
	"comments_service/graph/model"
	"comments_service/internal/storage"
	jwt "comments_service/pkg/JWT"
)


type SecureAccessComment struct{
	JWTKey string
}	

func NewAuthorization(JWTKey string) SecureAccessComment{
	return SecureAccessComment{ JWTKey: JWTKey}
}


func (auth SecureAccessComment) Authorize(login string) (string, error){
	token, err := jwt.CreateToken(auth.JWTKey, login)
	if err != nil{
		return "", err
	}
	return token, nil
}


func (auth SecureAccessComment) Authentication(IdentificationData model.IdentificationData, stor storage.Storage) error {
	loginFromToken, err := jwt.CheckUserToken(IdentificationData.Token, auth.JWTKey)
	if err != nil{
		return err
	}
	if loginFromToken != IdentificationData.Login{
		return custom_errors.ErrTokenOrUserInvalid
	}
	if err := stor.IsLoginExist(IdentificationData.Login); err != nil{
		return err
	}
	return nil
}

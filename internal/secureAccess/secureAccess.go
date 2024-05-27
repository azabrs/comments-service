package secure_access

import (
	custom_errors "comments_service/errors"
	"comments_service/graph/model"
	"comments_service/internal/storage"
	jwt "comments_service/pkg/JWT"
)


type SecureAccess struct{
	JWTKey string
}	

func NewAuthorization(JWTKey string) SecureAccess{
	return SecureAccess{ JWTKey: JWTKey}
}


func (auth SecureAccess) Authorize(login string) (string, error){
	token, err := jwt.CreateToken(auth.JWTKey, login)
	if err != nil{
		return "", err
	}
	return token, nil
}


func (auth SecureAccess) Authentication(IdentificationData model.IdentificationData, stor storage.Storage) error {
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

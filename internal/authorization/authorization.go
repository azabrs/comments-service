package authorization

import jwt "comments_service/pkg/JWT"


type Authorization struct{
	JWTKey string
}	

func NewAuthorization(JWTKey string) Authorization{
	return Authorization{ JWTKey: JWTKey}
}


func (auth Authorization) Authorize(login string) (string, error){
	token, err := jwt.CreateToken(auth.JWTKey, login)
	if err != nil{
		return "", err
	}
	return token, nil
}
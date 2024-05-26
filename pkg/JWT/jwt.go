package jwt

import (
	custom_errors "comments_service/errors"
	//"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct{
	*jwt.RegisteredClaims
	Login string `json:"Login"`
}

func CreateToken(signingKey string, login string) (string, error){
	claims := &Claims{
		
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: nil,//jwt.NewNumericDate(time.Now().Add(time.Hour * 1000000)),
		},
		Login : login,
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := Token.SignedString([]byte(signingKey))
	if err != nil{
		return "", err
	}
	return ss, nil
}

func CheckUserToken(tokenString string, signingKey string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil{
		return "", err
	}
	if !token.Valid{
		return "", custom_errors.ErrTokenOrUserInvalid
	}
	return claims.Login, nil
}
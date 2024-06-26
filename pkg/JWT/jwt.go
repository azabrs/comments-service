package jwt

import (
	custom_errors "comments_service/errors"
	"time"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct{
	*jwt.RegisteredClaims
	Login string `json:"Login"`
}

func CreateToken(signingKey string, login string) (string, error){
	claims := &Claims{
		
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1000000)),
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
	//claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil{
		return "", custom_errors.ErrTokenOrUserInvalid
	} else if claims, ok := token.Claims.(*Claims); ok && claims.Login != ""{
		return claims.Login, nil
	} else {
		return "", fmt.Errorf("unknown claims type, cannot proceed")
	}
}
package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct{
	jwt.RegisteredClaims
}

func CreateToken(signingKey string, expiresAt time.Time) (string, error){
	claims := Claims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := Token.SignedString(signingKey)
	if err != nil{
		return "", err
	}
	return ss, nil
}
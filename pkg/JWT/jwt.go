package jwt

import (

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct{
	jwt.RegisteredClaims
}

func CreateToken(signingKey string) (string, error){
	claims := Claims{
		jwt.RegisteredClaims{
			//ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := Token.SignedString(signingKey)
	if err != nil{
		return "", err
	}
	return ss, nil
}
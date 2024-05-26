package jwt

import (

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct{
	*jwt.RegisteredClaims
	Login string `json:"login"`
}

func CreateToken(signingKey string, login string) (string, error){
	claims := &Claims{
		
		RegisteredClaims: &jwt.RegisteredClaims{
			//ExpiresAt: jwt.NewNumericDate(expiresAt),
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
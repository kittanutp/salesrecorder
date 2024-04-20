package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kittanutp/salesrecorder/config"
)

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

func getSecretKey() string {
	return config.OAUTH2_SECRET
}

func GenerateToken(username string, isUser bool) string {
	claims := &authCustomClaims{
		username,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    username,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		panic(err)
	}
	return t
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("invalid token")

		}
		return []byte(getSecretKey()), nil
	})

}

package jwt

import (
	"easy-etcd/src/config"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var mySigningKey = []byte(config.AdminSecret)

// EncodeJWT JWT加密
func EncodeJWT(username string, password string) (string, error) {
	// Create the Claims
	claims := MyCustomClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

// DecodeJWT JWT解码
func DecodeJWT(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return "", "", err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims.Username, claims.Password, nil
	} else {
		return "", "", errors.New("无效Token")
	}
}

func UserJWT() {}

package jwtx

import (
	"github.com/golang-jwt/jwt"
)

// jwt 工具

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	signedString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

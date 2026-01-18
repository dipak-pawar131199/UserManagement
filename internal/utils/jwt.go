package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = "abcdeftehghdjdu"

func GenerateToken(userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userName,
		"exp":     time.Now().Add(24 * time.Hour).Unix()})
	jwtkeybyte := []byte(jwtKey)
	return token.SignedString(jwtkeybyte)
}

package blogapi

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(id int, username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"id":       int(id),
		"username": string(username),
		"iat":      time.Now().Unix(),
		"scopes":   []string{"api:read", "api:write"},
	})

	jwtToken, err := token.SignedString(Key)
	if err != nil {
		log.Println(err)
	}
	return jwtToken
}

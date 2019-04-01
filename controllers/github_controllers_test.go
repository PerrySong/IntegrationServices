package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"testing"
)

func TestGithubCallbackHandler(t *testing.T) {
	tokenString, _ := GenerateToken()
	fmt.Println(tokenString)

	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return "showcase_secret", nil
	})

	fmt.Println(token)

	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
}

func GenerateToken() (string, error) {
	var err error
	secret := "showcase_secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": "email",
		"id":    1314,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, nil
}

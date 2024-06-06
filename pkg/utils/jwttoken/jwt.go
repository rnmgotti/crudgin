package jwttoken

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Login string    `json:"login"`
	Exp   time.Time `json:"exp"`
	jwt.StandardClaims
}

func GenerateToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(time.Hour * 24),
	})

	tokenString, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		log.Println("[ERROR] Can't create token!", err)
		return "", err
	}

	return tokenString, nil
}

func ValidToken(tokenStr string) (login string, valid bool, err error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})
	if err != nil {
		log.Println("Error parsing token", err)
		return
	}

	if claims.Exp.Before(time.Now()) {
		valid = false
		log.Println("[WARNING] Token Expired")
	} else {
		valid = token.Valid
	}

	login = claims.Login

	return
}

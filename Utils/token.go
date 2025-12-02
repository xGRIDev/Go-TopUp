package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKeys = "hardestSecret"

func JWTGenerate(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKeys))
}

func VerifToken(token string) error {
	parsedTkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKeys), nil
	})
	if err != nil {
		fmt.Println("Could not parse token")
		return errors.New("could not parse the token")
	}

	tokenisValid := parsedTkn.Valid
	if !tokenisValid {
		return errors.New("invalid token.")
	}

	return nil
}

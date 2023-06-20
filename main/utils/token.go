package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sugandhasaxena1911/GoSamples/main/models"
	"log"
	"os"
)

func GenerateToken(user models.User) (string, error) {
	fmt.Println("Inside generate token ", user)
	secretstr := os.Getenv("SECRET")
	var err error
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	fmt.Println("value of token ", *token)

	tokenstr, err := token.SignedString([]byte(secretstr))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstr, nil
	//fmt.Println(token.s)
}

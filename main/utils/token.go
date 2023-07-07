package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sugandhasaxena1911/GoSamples/main/models"
	"os"
	"time"
)

func GenerateToken(user models.User) (string, error) {
	fmt.Println("Inside generate token ", user)
	secretstr := os.Getenv("SECRET")
	t := time.Now().Add(2 * time.Minute)
	fmt.Println("Time of expiration :  ", t)
	fmt.Println("Expiry :  ", jwt.NewNumericDate(t))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
		"exp":   t.Unix(),
	})
	t1, e := token.Claims.GetExpirationTime()
	fmt.Println("The value of expiry set is :   ", t1, e)

	fmt.Printf("value of token %#+v  ", *token)
	fmt.Printf("value of claims %#+v  ", token.Claims)

	//Raw is empty , signature empty , valid :false

	tokenstr, e := token.SignedString([]byte(secretstr))
	/*if e != nil {
		log.Fatalln(e)
	}*/ ///   checked in calling func
	return tokenstr, e
}

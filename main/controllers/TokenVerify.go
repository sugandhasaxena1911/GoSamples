package controllers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sugandhasaxena1911/GoSamples/main/models"
	"github.com/sugandhasaxena1911/GoSamples/main/utils"
	"net/http"
	"os"
	"strings"
)

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("TokenVerifyMiddleWare invoked")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err models.Error
		authHeader := r.Header.Get("Authorization") // Bearer hdjejefjegffdhfajhfkjfjfhjfdjhfdvjbhdv
		fmt.Println(authHeader)
		//string to slice
		bearertoken := strings.Split(authHeader, " ")
		fmt.Println(bearertoken)
		if len(bearertoken) == 2 {
			authtoken := bearertoken[1]
			//check if token valid
			token, error := jwt.Parse(authtoken, func(token *jwt.Token) (interface{}, error) {
				fmt.Printf("token before is   %+v ", token)
				// something related to assertion
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was error in token siging method")
				}
				return []byte(os.Getenv("SECRET")), nil
			})
			fmt.Printf("token after is   %+v ", token)
			fmt.Println("\nError  is", error)

			if error != nil {
				err.Message = error.Error()
				utils.RespondError(w, http.StatusUnauthorized, err)
				return
			}

			fmt.Println("TOKEN : ", token) // valid is true
			//t, _ := token.Claims.GetExpirationTime()
			fmt.Println("Token is VALID , claims is ", token.Claims)
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				err.Message = error.Error()
				utils.RespondError(w, http.StatusUnauthorized, err)
				return
			}
		} else {
			err.Message = "Token is not Valid "
			utils.RespondError(w, http.StatusUnauthorized, err)
			return
		}

	})

}

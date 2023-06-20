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
		var eobj models.Error
		authHeader := r.Header.Get("Authorization")
		fmt.Println(authHeader)
		//string to slice
		bearertoken := strings.Split(authHeader, " ")
		fmt.Println(bearertoken)
		if len(bearertoken) == 2 {
			authtoken := bearertoken[1]
			//check if token valid
			token, error := jwt.Parse(authtoken, func(token *jwt.Token) (interface{}, error) {
				fmt.Printf("token%+v ", token)
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was error in token")
				}
				return []byte(os.Getenv("SECRET")), nil
			})
			if error != nil {
				eobj.Message = error.Error()
				utils.RespondError(w, http.StatusUnauthorized, eobj)
				return
			}

			fmt.Println("TOKEN : ", token) // valid is true
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				eobj.Message = error.Error()
				utils.RespondError(w, http.StatusUnauthorized, eobj)
				return
			}
		} else {
			eobj.Message = "Token is not Valid "
			utils.RespondError(w, http.StatusUnauthorized, eobj)
			return
		}

	})

}

package domain

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const TOKENDURATION = time.Hour
const TOKENSIGNEDSTRING = "HELLOSIGNEDSTRING"

type User struct {
	Username    string         `json:"username"`
	Password    string         `json:"password,omitempty"`
	Role        string         `json:"role"`
	Customer_id sql.NullString `json:"customer_id"`
	Created_on  string         `json:"created_on"`
}

func (usr *User) GenerateHashPassword() ([]byte, error) {

	pass, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.MinCost)
	if err != nil {
		log.Println("Error generating Password hash ", err)
		return nil, err
	}
	return pass, nil
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserLoginResponse struct {
	Username    string         `json:"username"`
	Password    string         `json:"password,omitempty"`
	Role        string         `json:"role"`
	Customer_id sql.NullString `json:"customerId"`
	Accounts    sql.NullString `json:"accounts"`
}

func (usrloginres UserLoginResponse) GenerateToken() (*string, error) {
	// token will not have accounts if role=admin
	log.Println("Inside GenerateToken")
	var claims jwt.MapClaims
	log.Println("customer id  ", usrloginres.Customer_id.Valid)

	if usrloginres.Customer_id.Valid {
		claims = usrloginres.claimsForUser()
	} else {
		claims = usrloginres.claimsForAdmin()
	}

	log.Println("map claims ", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedtoken, err := token.SignedString([]byte(TOKENSIGNEDSTRING))
	log.Println("token ", token)
	log.Println("signedtoken ", signedtoken)

	if err != nil {
		return nil, err
	}

	return &signedtoken, nil

}
func (usrloginres UserLoginResponse) claimsForUser() jwt.MapClaims {
	log.Println("Inside claimsForUser")

	a := usrloginres.Accounts.String
	accnts := strings.Split(a[1:len(a)-1], ",")

	claims := jwt.MapClaims{
		"username":    usrloginres.Username,
		"role":        usrloginres.Role,
		"customer_id": usrloginres.Customer_id.String,
		"accounts":    accnts,
		"exp":         time.Now().Add(TOKENDURATION).Unix(),
	}
	return claims

}
func (usrloginres UserLoginResponse) claimsForAdmin() jwt.MapClaims {
	log.Println("Inside claimsForAdmin")

	claims := jwt.MapClaims{
		"username": usrloginres.Username,
		"role":     usrloginres.Role,
		"exp":      time.Now().Add(TOKENDURATION).Unix(),
	}
	return claims

}

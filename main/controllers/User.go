package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/sugandhasaxena1911/GoSamples/main/driver"
	"github.com/sugandhasaxena1911/GoSamples/main/models"
	"github.com/sugandhasaxena1911/GoSamples/main/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var db *sql.DB

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	e := json.NewDecoder(r.Body).Decode(&user)
	if e != nil {
		log.Fatalln(e)
	}
	log.Println(user)
	//empty email or password  : bad request
	err := models.Error{}

	if user.Email == "" {
		err.Message = "Email not provided for signup"
		utils.RespondError(w, http.StatusBadRequest, err)
		return

	}
	if user.Passwords == "" {
		err.Message = "Password not provided for signup"
		utils.RespondError(w, http.StatusBadRequest, err)
		return
	}

	//hash password : slice of byte
	hash, e := bcrypt.GenerateFromPassword([]byte(user.Passwords), bcrypt.DefaultCost)
	if e != nil {
		log.Fatalln(err)
	}
	fmt.Println("Encrypted password (Hash) ", hash)
	fmt.Println("password recieved ", user.Passwords)
	//set it to hash
	user.Passwords = string(hash)
	//now we need to store this user in db with hashed password
	db = driver.GetDBConnection()
	st := "insert into users (email,password) values ($1,$2) RETURNING id;"
	e = db.QueryRow(st, user.Email, user.Passwords).Scan(&user.ID)
	if e != nil {
		err.Message = fmt.Sprint("Internal server error while insertion ", e)
		utils.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	user.Passwords = ""
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login invoked")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	//validation on data
	//empty email or password  : bad request
	err := models.Error{}

	if user.Email == "" {
		err.Message = "Email not provided for login"
		utils.RespondError(w, http.StatusBadRequest, err)
		return

	}
	if user.Passwords == "" {
		err.Message = "Password not provided for login"
		utils.RespondError(w, http.StatusBadRequest, err)
		return
	}

	//check if data exists in the table
	password := user.Passwords
	fmt.Println(password)
	db = driver.GetDBConnection()
	st := "select * from users where email=$1"
	row := db.QueryRow(st, user.Email)
	e := row.Scan(&user.ID, &user.Email, &user.Passwords)
	if e != nil {
		if e == sql.ErrNoRows {
			err.Message = "user doesnt exists"
			utils.RespondError(w, http.StatusBadRequest, err)
			return
		} else {
			log.Fatalln(e)
		}
	}
	fmt.Println("the final user object ", user)

	// compare password passed in request and stored in db
	e = bcrypt.CompareHashAndPassword([]byte(user.Passwords), []byte(password))
	if e != nil {
		err.Message = "Password is incorrect . Please try again with correct password"
		utils.RespondError(w, http.StatusBadRequest, err)
		return
	}
	// All validation are passed now
	tokenstr, e := utils.GenerateToken(user)
	fmt.Println("token string is :     ", tokenstr)
	fmt.Println("error is :     ", e)
	if e != nil {
		log.Fatalln(e)
	}
	jsonwebtoken := models.JWT{tokenstr}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonwebtoken)

}
func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protectedEndpoint invoked")
	fmt.Println("HURRAYYYYYYYYYYYYYYYYYYYYYYY")

}

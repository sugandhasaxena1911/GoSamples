package domain

import (
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepositoryDB struct {
	client *sqlx.DB
}

func (userrepodb UsersRepositoryDB) RegisterUser(usr User) (*User, error) {

	log.Println("Inside RegisterUser ", usr)
	log.Println("Inside RegisterUser cust id ", usr.Customer_id)

	var id int
	//genearte hash
	pass, err := usr.GenerateHashPassword()
	if err != nil {
		return nil, errors.New("Unexpected error ")

	}

	sqlt := "insert into users_s(username ,password ,role ,customer_id ,created_on) values($1,$2,$3,$4,$5) RETURNING id"
	err = userrepodb.client.QueryRow(sqlt, usr.Username, string(pass), usr.Role, usr.Customer_id, usr.Created_on).Scan(&id)
	if err != nil {
		log.Println("Error1 ", err)

		return nil, errors.New("Unexpected error ")
	}
	log.Println("After RegisterUser ", usr, &usr)

	return &usr, nil

}

func (userrepodb UsersRepositoryDB) LoginUser(usrreq UserLoginRequest) (*UserLoginResponse, error) {
	var sqlt string
	sqlt = "select U.username,U.password,U.role , U.customer_id , array_agg(A.account_id) accounts from users_s U LEFT OUTER JOIN accounts A ON U.customer_id = A.customer_id where U.username = $1 group by U.username,U.password, U.role, U.customer_id"

	var usrloginres UserLoginResponse
	err := userrepodb.client.Get(&usrloginres, sqlt, usrreq.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Error occured ", err.Error())
			return nil, errors.New("Invalid Username")
		}
		log.Println("Error occured2 ", err.Error())

		return nil, errors.New("Unexpected DB Error")

	}

	err = bcrypt.CompareHashAndPassword([]byte(usrloginres.Password), []byte(usrreq.Password))
	if err != nil {
		log.Println("Error password hash", err.Error())

		return nil, errors.New("Incorrect Password")

	}
	log.Println("response ", usrloginres)

	return &usrloginres, nil

}

func NewUsersRepositoryDB(client *sqlx.DB) UsersRepositoryDB {
	return UsersRepositoryDB{client}
}

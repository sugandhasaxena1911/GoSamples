package driver

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/sugandhasaxena1911/GoSamples/main/models"
	"log"
	"os"
)

var db *sql.DB

func GetDBConnection() *sql.DB {
	//pgUrl, err := pq.ParseURL("postgres://eggcrzff:H4O_kCmg7oH3F1lG3KlybzGvMNbmhLs5@rajje.db.elephantsql.com/eggcrzff?sslmode=disable")
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	db, err = sql.Open("postgres", pgUrl)

	if err != nil {
		log.Fatalln(err)
	}
	e := db.Ping()
	if e == nil {
		fmt.Println("Connection established successfully ", db)
	}
	return db
}

func InsertUser(user *models.User) error {
	db = GetDBConnection()
	st := "insert into users (email,password) values ($1,$2) RETURNING id;"
	e := db.QueryRow(st, user.Email, user.Passwords).Scan(&user.ID)
	return e
}

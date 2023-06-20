package driver

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
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

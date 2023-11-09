package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/logger"
)

func NewSqlDBConnection() *sql.DB {
	var db *sql.DB
	pgUrl, err := pq.ParseURL("postgres://eggcrzff:H4O_kCmg7oH3F1lG3KlybzGvMNbmhLs5@rajje.db.elephantsql.com/eggcrzff")
	//pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	fmt.Println("CHECK")
	if err != nil {
		log.Fatalln("Cannot parse URl", err)
	}
	db, err = sql.Open("postgres", pgUrl)

	if err != nil {
		log.Fatalln(" Cannot open db connection ", err)
	}
	e := db.Ping()
	if e == nil {
		fmt.Println("Connection established successfully ", db)
	}
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db

}

func NewSqlxDBConnection() *sqlx.DB {
	var db *sqlx.DB
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDR")
	dbName := os.Getenv("DB_NAME")

	pgUrl, err := pq.ParseURL(fmt.Sprintf("%s://%s:%s@%s/%s", dbName, dbUser, dbPassword, dbAddress, dbUser))
	//pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logger.Info("Check is connection successful")
	if err != nil {
		logger.Fatal("Cannot parse URl" + err.Error())
	}
	db, err = sqlx.Open("postgres", pgUrl)

	if err != nil {
		logger.Fatal(" Cannot open db connection " + err.Error())
	}
	e := db.Ping()
	if e == nil {
		fmt.Println("Connection established successfully ", db)
	}
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db

}

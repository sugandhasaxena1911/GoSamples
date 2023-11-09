package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sugandhasaxena1911/GoSamples/BankingAuthApp/domain"
	"github.com/sugandhasaxena1911/GoSamples/BankingAuthApp/driver"
	"github.com/sugandhasaxena1911/GoSamples/BankingAuthApp/service"
)

var client *sqlx.DB

func init() {

	log.Println("inititalize DB connection ")
	client = driver.NewSqlxDBConnection()
}

func Start() {
	log.Println("inside start")

	r := mux.NewRouter()
	userhandler := UserHandler{service.NewDefaultUserservice(domain.NewUsersRepositoryDB(client))}
	r.HandleFunc("/auth/register", userhandler.Register).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", userhandler.Login).Methods(http.MethodPost)
	r.HandleFunc("/auth/logintoken", userhandler.LoginToken).Methods(http.MethodPost)

	log.Fatalln(http.ListenAndServe("localhost:8080", r))
	log.Println("after Listen & Serve")

}

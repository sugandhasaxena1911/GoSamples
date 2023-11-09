package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	mux "github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/domain"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/driver"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/logger"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/service"
)

var DBclient *sqlx.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Cannot load env variables ")
	}
}

func init() {

	DBclient = driver.NewSqlxDBConnection()

}

func SanityCheck() {

	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWORD") == "" ||
		os.Getenv("DB_ADDR") == "" ||
		os.Getenv("DB_NAME") == "" {
		log.Fatalln("Environment variables are not defined ")
	}

}

func Start() {
	/*
		http.HandleFunc("/greet", Greet)                     // register with default multiplexer  : DefaultServeMux
		http.HandleFunc("/GetAllCustomers", GetAllCustomers) // register with default multiplexer  : DefaultServeMux

		log.Fatalln(http.ListenAndServe(":8000", nil))

	*/
	// Using our own multiplexer, using gorilla mux
	SanityCheck()
	mux := mux.NewRouter()

	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	mux.HandleFunc("/customers/{customer_id:[A-Z0-9]+}", getCustomer).Methods(http.MethodGet)

	ch := CustomerHandler{service.NewCustomerServiceDefault(domain.NewCustomerRepositoryDB(DBclient))}
	mux.HandleFunc("/customersHexagonal", ch.getAllCustomersHexagonal).Methods(http.MethodGet)
	mux.HandleFunc("/customersHexagonal/{customer_id:[0-9]+}", ch.getCustomerByIdHexagonal).Methods(http.MethodGet)
	// account
	acchandler := AccountHandler{service.NewAcccountServiceDefault(domain.NewAccountRepositoryDB(DBclient))}
	mux.HandleFunc("/accounts", acchandler.CreateAccount).Methods(http.MethodPost)
	mux.HandleFunc("/accounts/{account_id:[0-9]+}", acchandler.GetAccount).Methods(http.MethodGet)

	mux.HandleFunc("/customersHexagonal/{customer_id:[0-9]+}/accounts", acchandler.CreateAccount).Methods(http.MethodPost)

	//txns
	th := Txnhandler{service.NewTxnservicedefault(domain.NewTxnRepositoryDB(DBclient))}
	mux.HandleFunc("/txns", th.PostTxns).Methods(http.MethodPost)
	/* PS : NewCustomerRepositoryDB() and NewAccountRepositoryDB() returns a struct that has a DB client ,
	means each time you need to take new db connection , which is not correct
	*/
	logger.Info("after handle func  ")
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), mux))
}

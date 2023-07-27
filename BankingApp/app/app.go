package app

import (
	"fmt"
	mux2 "github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	/*
		http.HandleFunc("/greet", Greet)                     // register with default multiplexer  : DefaultServeMux
		http.HandleFunc("/GetAllCustomers", GetAllCustomers) // register with default multiplexer  : DefaultServeMux

		log.Fatalln(http.ListenAndServe(":8000", nil))

	*/
	// Using our own multiplexer, using gorilla mux
	mux := mux2.NewRouter()

	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	mux.HandleFunc("/customers/{customer_id:[A-Z0-9]+}", getCustomer).Methods(http.MethodGet)

	fmt.Println("hello")
	log.Fatalln(http.ListenAndServe(":8000", mux))
}

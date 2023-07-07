package app

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", Greet)                     // register with default multiplexer  : DefaultServeMux
	http.HandleFunc("/GetAllCustomers", GetAllCustomers) // register with default multiplexer  : DefaultServeMux

	log.Fatalln(http.ListenAndServe(":8000", nil))
}

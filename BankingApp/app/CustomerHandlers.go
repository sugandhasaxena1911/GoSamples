package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/DTO"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/service"
)

type Customer struct {
	Name    string `json:"name" xml:"fullname"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type CustomerHandler struct {
	custservice service.CustomerService
}

func (ch *CustomerHandler) getAllCustomersHexagonal(w http.ResponseWriter, r *http.Request) {
	route := mux.CurrentRoute(r)
	log.Println("route is ", route.GetName())
	var custs []DTO.CustomerDto
	status := r.URL.Query().Get("status")
	fmt.Println("status ", status)
	custs, err := ch.custservice.GetAllCustomers(status)
	if err != nil {
		setResponse(w, err.Code, err.AsMessage())
		return
	}
	//if r.Header.Get("Content-Type") == "application/xml" {
	setResponse(w, http.StatusOK, custs)

}

func (ch *CustomerHandler) getCustomerByIdHexagonal(w http.ResponseWriter, r *http.Request) {
	mp := mux.Vars(r)
	id := mp["customer_id"]
	customer, err := ch.custservice.GetCustomerById(id)
	log.Println("Fetch the customer ", err)

	if err != nil {
		setResponse(w, err.Code, err.AsMessage())
		return
	}
	setResponse(w, http.StatusOK, customer)

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Fprint Hello	 World !")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside")

	customers := []Customer{
		{Name: "Sugandha ", Zipcode: "5660056", City: "Lucknow"},
		{Name: "Manu ", Zipcode: "11", City: "Dubai"},
		{Name: "Sapna ", Zipcode: "99056", City: "London"},
		{Name: "Saumya", Zipcode: "560", City: "Bangalore"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		xml.NewEncoder(w).Encode(&customers)

	} else { // default case:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&customers)
	}

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	//returns route variables
	fmt.Println("Inside get Customer")
	// Vars retuns a map with the request variables
	m1 := mux.Vars(r)
	fmt.Println(m1)

	custId := m1["customer_id"]
	fmt.Fprint(w, custId)
	fmt.Fprint(w, m1)

}
func createCustomer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Its a POST method ")
}

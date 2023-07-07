package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"fullname"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello	 World !")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Sugandha ", Zipcode: "5660056", City: "Lucknow"},
		{Name: "Manu ", Zipcode: "11", City: "Dubai"},
		{Name: "Hello ", Zipcode: "99056", City: "London"},
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

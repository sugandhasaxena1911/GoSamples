package GoSamples

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Accounts struct {
	AccName       string `json:"AccName"`
	AccId         int    `json:"AccId""`
	AccountHolder string `json:"AccHolder""`
}

var Accs []Accounts

func HandleAccounts() {

	r := mux.NewRouter()
	r.HandleFunc("/AccountsHome", AccountsHome).Methods("POST")
	r.HandleFunc("/AddAccount", AddAccount).Methods("POST")
	r.HandleFunc("/AddAccounts", AddAccounts).Methods("POST")

	log.Fatalln(http.ListenAndServe(":5000", r))
}

func AccountsHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello , Welcome to our Accounts Page")

}
func AddAccount(w http.ResponseWriter, r *http.Request) {
	acc := Accounts{}
	e := json.NewDecoder(r.Body).Decode(&acc)
	if e != nil {
		log.Println("Cannot decode ", e)
		return
	}
	Accs = append(Accs, acc)
	w.Header().Set("Content-Type", "application/json")
	e = json.NewEncoder(w).Encode(acc)
	if e != nil {
		log.Println("Cannot encode ", e)
	}

}

func AddAccounts(w http.ResponseWriter, r *http.Request) {
	accs := []Accounts{}
	e := json.NewDecoder(r.Body).Decode(&accs)
	if e != nil {
		log.Println("Cannot decode list of accounts ", e)
		return
	}
	fmt.Println(accs)
	// range over slice

	for _, v := range accs {
		Accs = append(Accs, v)

	}
	w.Header().Set("Content-Type", "application/json")
	e = json.NewEncoder(w).Encode(accs)
	if e != nil {
		log.Println("Cannot encode ", e)
	}

}

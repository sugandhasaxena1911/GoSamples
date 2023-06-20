package GoSamples

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	BookId   int    `json:"BookId"`
	BookName string `json:"BookName"`
}

var Books []Book

func HandleRequests() {
	r := mux.NewRouter() // Type : *mux.Router
	fmt.Printf("%T ", r)
	r.HandleFunc("/Home", HomePage).Methods("POST")
	r.HandleFunc("/Home2", HomePage2).Methods("POST")
	r.HandleFunc("/AddBook", AddBook).Methods("POST")
	r.HandleFunc("/AddBooks", AddBooks).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	w.Write([]byte("Welcome to my Library "))

}

func HomePage2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello to my Library -2")
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside")
	book1 := Book{}
	dec := json.NewDecoder(r.Body)
	e := dec.Decode(&book1)
	if e != nil {
		log.Fatalln(e)
	}
	Books := append(Books, book1)
	fmt.Println(book1)
	fmt.Println(Books)
	// draft the response: send the added book back
	w.Header().Set("Content-Type", "application/json")
	e = json.NewEncoder(w).Encode(book1)
	if e != nil {
		log.Fatalln(e)
	}
}

func AddBooks(w http.ResponseWriter, r *http.Request) {
	books := []Book{}
	json.NewDecoder(r.Body).Decode(&books)
	for _, v := range books {
		Books = append(Books, v)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)
}

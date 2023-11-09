package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello1")
	mux := mux.NewRouter()
	mux.HandleFunc("/mytest", mytest).Methods(http.MethodGet)
	mux.HandleFunc("/mytest/{one}/{two}", callme).Methods(http.MethodGet)
	log.Fatalln(http.ListenAndServe(":8001", mux))

	fmt.Println("Hello")
}
func mytest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "mytest")
}

func callme(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Callme")

	mp := mux.Vars(r)
	val := r.URL.Query()

	fmt.Println(mp["one"])
	fmt.Println(mp["two"])
	fmt.Println(val["three"])
	st := fmt.Sprint(mp["one"], mp["two"], r.URL.Query().Get("three"), val.Get("four"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, st)
}

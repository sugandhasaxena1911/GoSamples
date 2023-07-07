package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"github.com/sugandhasaxena1911/GoSamples/main/controllers"
	"github.com/sugandhasaxena1911/GoSamples/main/driver"
	"net/http"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.GetDBConnection()
	HandleAccountRequests()

}

func HandleAccountRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/protected", controllers.TokenVerifyMiddleWare(controllers.ProtectedEndpoint)).Methods("POST")
	http.ListenAndServe(":8000", router)
}

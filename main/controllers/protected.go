package controllers

import (
	"fmt"
	"net/http"
)

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protectedEndpoint invoked")
	fmt.Println("HURRAYYYYYYYYYYYYYYYYYYYYYYY")

}

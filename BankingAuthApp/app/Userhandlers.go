package app

import (
	"encoding/json"
	"log"
	"net/http"

	DTO "github.com/sugandhasaxena1911/GoSamples/BankingAuthApp/DTO"

	"github.com/sugandhasaxena1911/GoSamples/BankingAuthApp/service"
)

type UserHandler struct {
	usrservice service.Userservice
}

func (usrhandler UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside the register handler ")
	//var user domain.User
	var user DTO.RegisterUserDTOReq
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error in generating response1 ", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("Invalid Request")
		if err != nil {
			log.Println("Error in generating response2 ", err)
			log.Panic("Error in generating response ", err)
		}
		return
	}
	usr, err := usrhandler.usrservice.RegisterUser(user)
	log.Println("The response DTO ", usr)

	if err != nil {
		log.Println("Error in generating response3 ", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Panic("Error in generating response ", err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(usr)
	if err != nil {
		log.Panic("Error in generating response ", err)
	}

}

func (usrhandler UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside the login handler ")
	var userloginreq DTO.LoginUserDTOReq
	err := json.NewDecoder(r.Body).Decode(&userloginreq)
	if err != nil {
		log.Println("Error in generating response1 ", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("Invalid Request")
		if err != nil {
			log.Println("Error in generating response2 ", err)
			log.Panic("Error in generating response ", err)
		}
		return
	}
	usrres, err := usrhandler.usrservice.LoginUser(userloginreq)
	if err != nil {
		log.Println("Error in generating response3 ", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(err.Error())
		if err != nil {
			log.Panic("Error in generating response ", err.Error())
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(usrres)
	if err != nil {
		log.Panic("Error in generating response ", err)
	}
}

func (usrhandler UserHandler) LoginToken(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside the login handler ")
	var userloginreq DTO.LoginUserDTOReq
	err := json.NewDecoder(r.Body).Decode(&userloginreq)
	if err != nil {
		log.Println("Error in generating response1 ", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("Invalid Request")
		if err != nil {
			log.Println("Error in generating response2 ", err)
			log.Panic("Error in generating response ", err)
		}
		return
	}
	tokenstring, err := usrhandler.usrservice.LoginUserToken(userloginreq)
	if err != nil {
		log.Println("Error in generating response3 ", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(err.Error())
		if err != nil {
			log.Panic("Error in generating response ", err.Error())
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tokenstring)
	if err != nil {
		log.Panic("Error in generating response ", err)
	}
}

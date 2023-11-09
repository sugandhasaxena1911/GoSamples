package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/DTO"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/logger"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/service"
)

type AccountHandler struct {
	accservice service.AccountService
}

func (acchandler AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	m1 := mux.Vars(r)
	if _, ok := m1["account_id"]; !ok {
		setResponse(w, http.StatusNotFound, "Path not found")

	}
	logger.Info("The account id is " + fmt.Sprintln(m1["account_id"]))
	acc, err := acchandler.accservice.GetAccount(m1["account_id"])
	if err != nil {
		setResponse(w, err.Code, err.AsMessage())
		return
	}
	setResponse(w, http.StatusOK, acc)
}

func (acchandler AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	logger.Info("Inside acc handler create account   ")

	var account DTO.AccountDtoRequest
	e := json.NewDecoder(r.Body).Decode(&account)
	if e != nil {
		setResponse(w, http.StatusBadRequest, e.Error())
		return
	}

	m1 := mux.Vars(r)
	if v, ok := m1["customer_id"]; ok {
		account.Customer_id = v

	}

	logger.Info("checking request recieved  " + fmt.Sprintln(account))
	acc, err := acchandler.accservice.CreateAccount(account)
	if err != nil {
		setResponse(w, err.Code, err.AsMessage())
		return
	}
	setResponse(w, http.StatusCreated, acc)

}

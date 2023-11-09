package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sugandhasaxena1911/GoSamples/BankingApp/DTO"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/logger"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/service"
)

type Txnhandler struct {
	txnservice service.Txnservice
}

func (txnhandler Txnhandler) PostTxns(w http.ResponseWriter, r *http.Request) {

	var txnreq DTO.TxnDTORequest
	err := json.NewDecoder(r.Body).Decode(&txnreq)
	if err != nil {
		setResponse(w, http.StatusBadRequest, err.Error())
		return

	}
	logger.Info("the txn request recieved " + fmt.Sprintln(txnreq))

	txnres, er := txnhandler.txnservice.PostTxns(txnreq)
	logger.Info("the txn response  recieved " + fmt.Sprintln(txnres))

	if er != nil {
		setResponse(w, er.Code, er.AsMessage())
		return
	}

	setResponse(w, http.StatusOK, txnres)
}

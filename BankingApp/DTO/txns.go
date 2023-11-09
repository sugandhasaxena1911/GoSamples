package DTO

import (
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/CONSTANTS"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/errs"
)

type TxnDTORequest struct {
	Account_id string `json:"account_id"`
	Amount     int    `json:"txn_amount"`
	Txn_type   string `json:"txn_type"`
	Txn_date   string `json:"txn_date"`
}

func (txn TxnDTORequest) ValidateTxnType() *errs.AppError {
	if txn.Txn_type != CONSTANTS.Withdrawal && txn.Txn_type != CONSTANTS.Deposit {
		return errs.NewBadError("Invalid txn type ")

	}
	return nil
}

func (txn TxnDTORequest) IsTxnTypeWithdrawal() bool {
	if txn.Txn_type != CONSTANTS.Withdrawal {
		return true
	} else {
		return false

	}
}

type TxnDTOResponse struct {
	Txn_id  string
	Balance int
}

func (txnreq TxnDTORequest) TxnTypeAsCode() string {
	var txntype string
	switch txnreq.Txn_type {
	case CONSTANTS.Withdrawal:
		txntype = "w"
	case CONSTANTS.Deposit:
		txntype = "d"

	}
	return txntype
}

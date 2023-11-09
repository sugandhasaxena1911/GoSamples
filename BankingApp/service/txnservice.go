package service

import (
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/DTO"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/domain"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/errs"
)

type Txnservice interface {
	PostTxns(txndtoreq DTO.TxnDTORequest) (*DTO.TxnDTOResponse, *errs.AppError)
}

type Txnservicedefault struct {
	txnrepo domain.TxnRepository
}

func (txndefault Txnservicedefault) PostTxns(txndtoreq DTO.TxnDTORequest) (*DTO.TxnDTOResponse, *errs.AppError) {
	err := txndefault.validateTxnRequest(txndtoreq)
	if err != nil {
		return nil, err
	}
	txndomain := domain.Transaction{Txn_id: "",
		Account_id: txndtoreq.Account_id,
		Amount:     txndtoreq.Amount,
		Txn_type:   txndtoreq.TxnTypeAsCode(),
		Txn_date:   txndtoreq.Txn_date}

	txn, err := txndefault.txnrepo.PostTxns(txndomain)
	if err != nil {

		return nil, err
	}
	return &DTO.TxnDTOResponse{Txn_id: txn.Txn_id, Balance: txn.Amount}, nil
}

func (txndefault Txnservicedefault) validateTxnRequest(txndtoreq DTO.TxnDTORequest) *errs.AppError {

	// check for account id exist or not

	// validate txn type
	er := txndtoreq.ValidateTxnType()
	if er != nil {
		return er
	}

	if txndtoreq.Amount < 0 {
		return errs.NewBadError("please enter valid amount")
	}

	return nil
}

func NewTxnservicedefault(txnrepo domain.TxnRepository) Txnservicedefault {
	return Txnservicedefault{txnrepo}
}

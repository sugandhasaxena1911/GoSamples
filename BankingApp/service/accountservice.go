package service

import (
	"fmt"

	"github.com/sugandhasaxena1911/GoSamples/BankingApp/DTO"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/domain"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/errs"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/logger"
)

type AccountService interface {
	CreateAccount(accdtoreq DTO.AccountDtoRequest) (*DTO.AccountDtoResponse, *errs.AppError)
	GetAccount(account_id string) (*DTO.AccountDtoResponse, *errs.AppError)
}

type AcccountServiceDefault struct {
	accdb domain.AccountrepositoryDB
}

func (accservicedef AcccountServiceDefault) GetAccount(account_id string) (*DTO.AccountDtoResponse, *errs.AppError) {
	logger.Info("service The account id is " + account_id)

	acc, err := accservicedef.accdb.GetAccount(account_id)
	if err != nil {
		return nil, err
	}
	account := acc.ToAccountDTOResponse()
	logger.Info("The account response  is " + fmt.Sprintln(account))
	return &account, err
}

func (accdefault AcccountServiceDefault) CreateAccount(accdtoreq DTO.AccountDtoRequest) (*DTO.AccountDtoResponse, *errs.AppError) {
	logger.Info("Inside account service")
	err := accdtoreq.ValidateAccDTORequest()
	if err != nil {
		return nil, err
	}
	acctype, _ := accdtoreq.AccTypeAsCode()
	accstatus, _ := accdtoreq.AccStatusAsCode()

	accdomain := domain.Account{Account_id: "", Customer_id: accdtoreq.Customer_id,
		Opening_date: accdtoreq.Opening_date, Account_type: acctype, Amount: accdtoreq.Amount,
		Status: accstatus}
	acc, err := accdefault.accdb.CreateAccount(accdomain)

	if err != nil {
		return nil, err
	}
	accdto := acc.ToAccountDTOResponse()
	return &accdto, nil

}

func NewAcccountServiceDefault(accdb domain.AccountrepositoryDB) AcccountServiceDefault {
	return AcccountServiceDefault{accdb}

}

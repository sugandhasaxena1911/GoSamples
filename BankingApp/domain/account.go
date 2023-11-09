package domain

import (
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/DTO"
)

type Account struct {
	Account_id   string
	Customer_id  string
	Opening_date string
	Account_type string
	Amount       float64
	Status       string
}

func (c Account) IsBalanceSufficient(amount float64) bool {
	if c.Amount < amount {
		return false
	}
	return true
}

func (c Account) ToAccountDTOResponse() DTO.AccountDtoResponse {
	return DTO.AccountDtoResponse{
		Account_id:   c.Account_id,
		Customer_id:  c.Customer_id,
		Opening_date: c.Opening_date,
		Account_type: c.AccountTypeAsText(),
		Amount:       c.Amount,
		Status:       c.AccountStatusAsText(),
	}
}

func (c Account) AccountTypeAsText() string {
	var acctype string
	switch c.Account_type {
	case "s":
		acctype = "savings"
	case "c":
		acctype = "current"
	case "f":
		acctype = "fixed"
	}

	return acctype
}

func (c Account) AccountStatusAsText() string {
	var status string
	switch c.Status {
	case "1":
		status = "active"
	case "2":
		status = "inactive"
	case "3":
		status = "dormant"
	}

	return status
}

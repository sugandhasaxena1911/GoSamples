package DTO

import "github.com/sugandhasaxena1911/GoSamples/BankingApp/errs"

//import "github.com/sugandhasaxena1911/GoSamples/BankingApp/domain"

type AccountDtoRequest struct {
	Customer_id  string  `json:"customer_id"`
	Opening_date string  `json:"opening_date"`
	Account_type string  `json:"account_type"`
	Amount       float64 `json:"amount"`
	Status       string  `json:"status"`
}

type AccountDtoResponse struct {
	Account_id   string  `json:"account_id"`
	Customer_id  string  `json:"customer_id"`
	Opening_date string  `json:"opening_date"`
	Account_type string  `json:"account_type"`
	Amount       float64 `json:"amount"`
	Status       string  `json:"status"`
}

/*
type Cyclic interface {
	ToAccountDTOResponse() AccountDtoResponse
	AccountTypeAsText() string
	AccountStatusAsText() string
}
*/
/*
func (accdtoreq AccountDtoRequest) ToAccountDomain() Cyclic {
	return domain.Account{Account_id: "", Customer_id: accdtoreq.Customer_id,
		Opening_date: accdtoreq.Opening_date, Account_type: accdtoreq.Account_type, Amount: accdtoreq.Amount,
		Status: accdtoreq.Status}

}
*/

func (accdtoreq AccountDtoRequest) AccTypeAsCode() (string, *errs.AppError) {
	var acctype string
	switch accdtoreq.Account_type {
	case "savings":
		acctype = "s"
	case "current":
		acctype = "c"
	case "fixed":
		acctype = "f"
	default:
		acctype = ""

	}
	if acctype == "" {
		return "", errs.NewBadError("Invalid Account Type")
	}
	return acctype, nil

}

func (accdtoreq AccountDtoRequest) AccStatusAsCode() (string, *errs.AppError) {
	var accstatus string
	switch accdtoreq.Status {
	case "active":
		accstatus = "1"
	case "inactive":
		accstatus = "2"
	case "dormant":
		accstatus = "3"
	default:
		accstatus = ""

	}
	if accstatus == "" {
		return "", errs.NewBadError("Invalid Account Status")
	}
	return accstatus, nil

}

func (accdtoreq AccountDtoRequest) ValidateAccDTORequest() *errs.AppError {
	_, err := accdtoreq.AccTypeAsCode()
	if err != nil {
		return err
	}
	_, err = accdtoreq.AccStatusAsCode()
	if err != nil {
		return err
	}

	if accdtoreq.Amount < 5000 {
		return errs.NewBadError("Insufficient amount. Please provide value >5000")
	}
	return nil
}

type LoginAccountDTO struct {
	Account_id string `json:"account_id"`
}

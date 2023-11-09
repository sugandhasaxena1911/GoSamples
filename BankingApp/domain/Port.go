package domain

import "github.com/sugandhasaxena1911/GoSamples/BankingApp/errs"

type CustomerRepository interface {
	FindAllCustomers(status string) ([]Customer, *errs.AppError)
	FindCustomerById(id string) (*Customer, *errs.AppError)
}

type AccountRepository interface {
	CreateAccount(account Account) (*Account, *errs.AppError)
}

type TxnRepository interface {
	PostTxns(txn Transaction) (*Transaction, *errs.AppError)
}

package domain

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/errs"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/logger"
)

type AccountrepositoryDB struct {
	client *sqlx.DB
}

func (accrepdb AccountrepositoryDB) GetAccount(account_id string) (*Account, *errs.AppError) {
	logger.Info("domian The account id is " + account_id)

	var account Account
	sqlst := "select account_id,customer_id,opening_date,account_type,amount,status from accounts where account_id = $1"
	err := accrepdb.client.Get(&account, sqlst, account_id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Info("Error is " + err.Error())

			return nil, errs.NewBadError("Invalid Account id")
		}
		logger.Info("Error " + err.Error())

		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}
	return &account, nil
}

func (acc AccountrepositoryDB) CreateAccount(account Account) (*Account, *errs.AppError) {
	logger.Info("Inside account repository db ")
	logger.Info(fmt.Sprintln("account ", account))

	var id int
	insertst := "insert into accounts(customer_id,opening_date,account_type,amount,status) values($1,$2,$3,$4,$5) RETURNING account_id;"
	//result, err := acc.client.Exec(insertst, account.Customer_id, account.Opening_date, account.Account_type, account.Amount, account.status)
	err := acc.client.QueryRow(insertst, account.Customer_id, account.Opening_date, account.Account_type, account.Amount, account.Status).Scan(&id)
	if err != nil {
		logger.Info("Error while inserting the account details" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	logger.Info(fmt.Sprintln("acc id ", id))

	account.Account_id = strconv.FormatInt(int64(id), 10)
	logger.Info(fmt.Sprintln("account ", account))

	return &account, nil
}

func NewAccountRepositoryDB(DBclient *sqlx.DB) AccountrepositoryDB {
	return AccountrepositoryDB{DBclient}
}

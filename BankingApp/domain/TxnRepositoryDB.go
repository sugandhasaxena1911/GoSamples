package domain

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/errs"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/logger"
)

type TxnRepositoryDB struct {
	client *sqlx.DB
}

func (txnrepodb TxnRepositoryDB) PostTxns(txns Transaction) (*Transaction, *errs.AppError) {
	logger.Info("Inside domain post txns")
	//check for the validations
	e := txnrepodb.validatePostTxn(txns)
	if e != nil {
		return nil, e
	}

	// start txn BEGIN
	logger.Info("initiate the transaction ")
	tx, err := txnrepodb.client.Begin()
	if err != nil {
		logger.Info("cannot begin the transaction ")
		return nil, errs.NewUnexpectedError("Unexpected database error ")

	}
	// insert into transation table
	logger.Info("insert into transation table")

	var id int64
	sqlst := "insert into transactions(account_id,amount,txn_type,txn_date) values ($1,$2,$3,$4) RETURNING txn_id"
	err = txnrepodb.client.QueryRow(sqlst, txns.Account_id, txns.Amount, txns.Txn_type, txns.Txn_date).Scan(&id)
	logger.Info("the id returned is " + fmt.Sprintln(id))
	if err != nil {
		logger.Info("issue " + fmt.Sprintln(err))
		return nil, errs.NewUnexpectedError("Some network problem : DB Issue")
	}

	// update the accounts table
	if txns.Txn_type == "w" {
		_, err = tx.Exec("Update accounts set amount = amount - $1 where account_id = $2", txns.Amount, txns.Account_id)
	} else {
		_, err = tx.Exec("Update accounts set amount = amount + $1 where account_id = $2", txns.Amount, txns.Account_id)

	}
	if err != nil {
		tx.Rollback()
		logger.Info("cannot update accounts id with balance  " + fmt.Sprintln(err))
		return nil, errs.NewUnexpectedError("Some network problem : DB Issue")
	}
	// get updated balance in accounts
	var balance int
	sqlst = "select amount from accounts where account_id = $1"
	//row := custs.client.QueryRow(sqlst, id)
	err = txnrepodb.client.Get(&balance, sqlst, txns.Account_id)
	if err != nil {
		tx.Rollback()
		logger.Info("Unable to fetch updated balace  " + fmt.Sprintln(err))
		return nil, errs.NewUnexpectedError("Some network problem : DB Issue")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Info("cannot commit the transaction ")
		return nil, errs.NewUnexpectedError("Unexpected database error ")

	}

	// start getting id since everything is committed
	logger.Info("the id returned is " + fmt.Sprintln(id))
	txns.Txn_id = strconv.FormatInt(int64(id), 10)
	txns.Amount = balance

	// start txn ends
	return &txns, nil

}

func (txnrepodb TxnRepositoryDB) validatePostTxn(txns Transaction) *errs.AppError {
	//check if account id exists
	var amount int
	sqlst := "select amount from accounts where account_id = $1"
	err := txnrepodb.client.Get(&amount, sqlst, txns.Account_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewBadError("Account Id does not exist.Please provide valid account id")
		} else {
			return errs.NewUnexpectedError("Unexpected DB error ")
		}
	}
	// check if sufficient balance for withdrawal
	if txns.Txn_type == "w" && amount < txns.Amount {
		return errs.NewBadError("Insufficient account balance")
	}

	return nil
}
func NewTxnRepositoryDB(client *sqlx.DB) TxnRepositoryDB {
	return TxnRepositoryDB{client}

}

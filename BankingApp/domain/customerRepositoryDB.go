package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/errs"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/logger"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (custs CustomerRepositoryDB) FindAllCustomers(status string) ([]Customer, *errs.AppError) {
	logger.Info("Inside FindAllCustomers")

	var findallsql string
	//var rows *sql.Rows
	customers := []Customer{}

	var err error
	if status == "Invalid" {
		logger.Error("Invalid status code")
		return nil, errs.NewNotFoundError("Invalid status code")

	}
	if status != "" {
		findallsql = "select customer_id, name, city , zipcode,date_of_birth, status from customers where status= $1"
		//rows, err = custs.client.Query(findallsql, status)
		err = custs.client.Select(&customers, findallsql, status)

	} else {
		//findallsql = "select customer_id, name, city , zipcode,date_of_birth, status from customers where $1='' "
		findallsql = "select customer_id, name, city , zipcode,date_of_birth, status from customers "
		//rows, err = custs.client.Query(findallsql)
		err = custs.client.Select(&customers, findallsql)

	}

	if err != nil {
		logger.Error("Error occured in query" + err.Error())

		/*
			// The Query will not treat no result set as error , only QueryRow expects atleast one row & hence no result --> error
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("No Customer Found.")
			} else {

				return nil, errs.NewUnexpectedError("Unexpected Database Error")
			}
		*/
		return nil, errs.NewUnexpectedError("Unexpected Database Error")

	}
	/*
		//Using the sqlx
		err = sqlx.StructScan(rows, &customers)
		if err != nil {
			logger.Error("Error occured in scanning record. " + err.Error())
			return nil, errs.NewUnexpectedError("Error in fetching the customer list")
		}
	*/

	/*
		count := 0
		for rows.Next() {
			var c Customer
			err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
			if err != nil {
				logger.Error("Error occured in scanning record" + err.Error())
				return nil, errs.NewUnexpectedError("Error in fetching the customer from the list")
			}
			customers = append(customers, c)
			count++
		}
		if count == 0 {
			logger.Error("No rows found ")
			return nil, errs.NewNotFoundError("No Customer Found.")

		}
	*/
	if len(customers) == 0 {
		logger.Error("No rows found ")
		return nil, errs.NewNotFoundError("No Customer Found.")
	}
	return customers, nil
}

func NewCustomerRepositoryDB(DBclient *sqlx.DB) CustomerRepositoryDB {
	//DBclient := driver.NewSqlDBConnection()
	//DBclient := driver.NewSqlxDBConnection()
	return CustomerRepositoryDB{DBclient}
}

// Find Customer by ID
func (custs CustomerRepositoryDB) FindCustomerById(id string) (*Customer, *errs.AppError) {
	var c Customer
	sqlst := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = $1"
	//row := custs.client.QueryRow(sqlst, id)
	err := custs.client.Get(&c, sqlst, id)

	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		logger.Error("Unable to fetch the record " + err.Error())

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not found.")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error.")
		}

	}

	return &c, nil
}

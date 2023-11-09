package domain

import (
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/DTO"
)

type Customer struct {
	Id          string `json:"id" db:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
	Status      string `json:"status"`
}

func (c Customer) StatusAsText() string {
	var status string
	switch c.Status {
	case "1":
		status = "Active"
	case "0":
		status = "Inactive"

	}
	return status
}
func (c Customer) ToCustDTO() DTO.CustomerDto {

	return DTO.CustomerDto{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.StatusAsText(),
	}

}

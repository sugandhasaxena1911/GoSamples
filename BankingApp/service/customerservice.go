package service

import (
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/DTO"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/domain"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]DTO.CustomerDto, *errs.AppError)
	GetCustomerById(string) (*DTO.CustomerDto, *errs.AppError)
}

type CustomerServiceDefault struct {
	custStub domain.CustomerRepository
}

func (service CustomerServiceDefault) GetAllCustomers(status string) ([]DTO.CustomerDto, *errs.AppError) {
	switch status {
	case "Active":
		status = "1"
	case "Inactive":
		status = "0"
	case "":
		status = ""
	default:
		status = "Invalid"

	}

	customers := []DTO.CustomerDto{}
	custs, err := service.custStub.FindAllCustomers(status)
	if err != nil {
		return nil, err
	}

	for _, c := range custs {
		customers = append(customers, c.ToCustDTO())
	}

	return customers, nil

}

func NewCustomerServiceDefault(custstub domain.CustomerRepository) CustomerServiceDefault {
	return CustomerServiceDefault{custstub}
}

func (service CustomerServiceDefault) GetCustomerById(id string) (*DTO.CustomerDto, *errs.AppError) {
	c, err := service.custStub.FindCustomerById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToCustDTO()
	return &response, nil
}

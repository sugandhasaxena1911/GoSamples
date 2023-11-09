package DTO

import (
	"database/sql"

	"github.com/sugandhasaxena1911/GoSamples/BankingAuthApp/domain"
)

type LoginAccountDTO struct {
	Account_id string `json:"account_id"`
}

type TokenDTO struct {
	Token *string
}
type RegisterUserDTOReq struct {
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"`
	Role        string `json:"role"`
	Customer_id string `json:"customer_id"`
	Created_on  string `json:"created_on"`
}

type RegisterUserDTORes struct {
	Username    string `json:"username"`
	Role        string `json:"role"`
	Customer_id string `json:"customer_id"`
	Created_on  string `json:"created_on"`
}

type LoginUserDTOReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserDTORes struct {
	Username    string            `json:"username"`
	Role        string            `json:"role"`
	Customer_id string            `json:"customer_id"`
	Accounts    []LoginAccountDTO `json:"Accounts,omitempty"`
}

func (registerReqDTO RegisterUserDTOReq) ToDomainUser() domain.User {
	return domain.User{Username: registerReqDTO.Username,
		Password:    registerReqDTO.Password,
		Role:        registerReqDTO.Role,
		Customer_id: ToNullString(registerReqDTO.Customer_id),
		Created_on:  registerReqDTO.Created_on}
}

func ToNullString(str string) sql.NullString {
	if str == "" {
		return sql.NullString{String: "", Valid: false}
	} else {
		return sql.NullString{String: str, Valid: true}
	}
}

package service

import (
	"log"
	"strings"

	DTO "github.com/sugandhasaxena1911/GoSamples/BankingAuthApp/DTO"
	"github.com/sugandhasaxena1911/GoSamples/BankingAuthApp/domain"
)

type Userservice interface {
	RegisterUser(usr DTO.RegisterUserDTOReq) (*DTO.RegisterUserDTORes, error)
	LoginUserToken(usr DTO.LoginUserDTOReq) (DTO.TokenDTO, error)
	LoginUser(usrreq DTO.LoginUserDTOReq) (*DTO.LoginUserDTORes, error)
}

type DefaultUserservice struct {
	userrepo domain.UsersRepository
}

func (defusrservice DefaultUserservice) LoginUserToken(usrreq DTO.LoginUserDTOReq) (DTO.TokenDTO, error) {
	usrloginres, err := defusrservice.userrepo.LoginUser(domain.UserLoginRequest{Username: usrreq.Username, Password: usrreq.Password})
	if err != nil {
		log.Println("error in service ", err.Error())
		return DTO.TokenDTO{}, err
	}
	tokenstring, err := usrloginres.GenerateToken()
	log.Println("tokenstring ", tokenstring)

	if err != nil {
		return DTO.TokenDTO{}, err
	}
	return DTO.TokenDTO{Token: tokenstring}, nil

}

func (defusrservice DefaultUserservice) RegisterUser(usr DTO.RegisterUserDTOReq) (*DTO.RegisterUserDTORes, error) {
	log.Println("The request DTO ", usr)
	user, err := defusrservice.userrepo.RegisterUser(usr.ToDomainUser())
	log.Println("The response DB ", user)

	if err != nil {

		return nil, err
	}
	var custid string
	if user.Customer_id.Valid {
		custid = user.Customer_id.String
	} else {
		custid = ""
	}
	return &DTO.RegisterUserDTORes{Username: user.Username, Role: user.Role, Customer_id: custid, Created_on: user.Created_on}, nil
}

func (defusrservice DefaultUserservice) LoginUser(usrreq DTO.LoginUserDTOReq) (*DTO.LoginUserDTORes, error) {
	usrloginres, err := defusrservice.userrepo.LoginUser(domain.UserLoginRequest{Username: usrreq.Username, Password: usrreq.Password})
	if err != nil {
		log.Println("error in service ", err.Error())
		return nil, err
	}

	//handle customer id
	var custid string
	if usrloginres.Customer_id.Valid {
		custid = usrloginres.Customer_id.String
	} else {
		custid = ""
	}
	//handle acconts
	var accdtos []DTO.LoginAccountDTO
	if usrloginres.Accounts.Valid {
		log.Println("Accounts ", usrloginres.Accounts.String)
		a := usrloginres.Accounts.String
		acc := a[1 : len(a)-1]
		accnts := strings.Split(acc, ",")
		for _, v := range accnts {
			log.Println("Account id ", v)
			d := DTO.LoginAccountDTO{Account_id: v}
			accdtos = append(accdtos, d)
		}

	} else {
		accdtos = nil
	}

	return &DTO.LoginUserDTORes{Username: usrloginres.Username, Role: usrloginres.Role, Customer_id: custid, Accounts: accdtos}, nil

}

func NewDefaultUserservice(usrrepodb domain.UsersRepositoryDB) DefaultUserservice {
	return DefaultUserservice{usrrepodb}

}

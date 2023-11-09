package domain

type UsersRepository interface {
	RegisterUser(usr User) (*User, error)
	LoginUser(usrreq UserLoginRequest) (*UserLoginResponse, error)
}

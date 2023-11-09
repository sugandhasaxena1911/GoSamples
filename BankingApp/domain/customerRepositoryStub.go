package domain

type CustomerRepositoryStub struct {
	custs []Customer
}

func (c CustomerRepositoryStub) FindAllCustomers() ([]Customer, error) {
	return c.custs, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	// create dummy values

	customers := []Customer{
		Customer{Id: "1", Name: "Sugandha", City: "Lucknow", Zipcode: "560035", Status: "Active"},
		Customer{Id: "2", Name: "Sapna", City: "Lucknow", Zipcode: "560035", Status: "Active"},
		Customer{Id: "3", Name: "Saumya", City: "Lucknow", Zipcode: "560035", Status: "Active"},
	}

	return CustomerRepositoryStub{customers}

}

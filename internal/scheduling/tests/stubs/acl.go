package stubs

import "github.com/zafir-co-ao/onna-narciso/internal/scheduling"

var Pacl scheduling.ProfessionalACLFunc = func(id string) (scheduling.Professional, error) {
	switch id {
	case "1":
		return scheduling.Professional{ID: "1", Name: "Sara Gomes"}, nil
	case "2":
		return scheduling.Professional{ID: "2", Name: "Julieta Venegas"}, nil
	case "3":
		return scheduling.Professional{ID: "3", Name: "John Doe"}, nil
	default:
		return scheduling.Professional{}, scheduling.ErrProfessionalNotFound
	}
}

var Sacl scheduling.ServiceACLFunc = func(id string) (scheduling.Service, error) {
	switch id {
	case "1":
		return scheduling.Service{ID: "1", Name: "Manicure"}, nil
	case "2":
		return scheduling.Service{ID: "2", Name: "Pedicure"}, nil
	case "3":
		return scheduling.Service{ID: "3", Name: "Depilação"}, nil
	case "4":
		return scheduling.Service{ID: "4", Name: "Manicure + Pedicure"}, nil
	default:
		return scheduling.Service{}, scheduling.ErrServiceNotFound
	}
}

type CustomerACLStub struct{}

func (c CustomerACLStub) FindCustomerByID(id string) (scheduling.Customer, error) {
	switch id {
	case "1":
		return scheduling.Customer{ID: "1", Name: "João Silva"}, nil
	case "2":
		return scheduling.Customer{ID: "2", Name: "Maria Oliveira"}, nil
	case "3":
		return scheduling.Customer{ID: "3", Name: "Carlos Ferreira"}, nil
	default:
		return scheduling.Customer{}, scheduling.ErrCustomerNotFound
	}
}

func (c CustomerACLStub) RequestCustomerRegistration(name string, phone string) (scheduling.Customer, error) {
	if name == "" || phone == "" {
		return scheduling.Customer{}, scheduling.ErrCustomerRegistration
	}

	return scheduling.Customer{ID: "1", Name: scheduling.Name(name), PhoneNumber: phone}, nil
}

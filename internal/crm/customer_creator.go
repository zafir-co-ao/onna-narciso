package crm

type CustomerCreator interface {
	Create() error
}

type customerCreatorImpl struct{}

func NewCustomerCreator() CustomerCreator {
	return &customerCreatorImpl{}
}

func (u *customerCreatorImpl) Create() error {
	return nil
}

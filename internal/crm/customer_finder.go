package crm

type CustomerFinder interface {
	Find() ([]CustomerOutput, error)
}

func NewCustomerFinder() customerFinderImpl {
	return customerFinderImpl{}
}

type customerFinderImpl struct {}

func (u customerFinderImpl) Find() ([]CustomerOutput, error) {
	return nil, nil
}
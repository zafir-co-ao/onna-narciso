package scheduling

type CustomerRepository interface {
	Get(id string) (Customer, error)
	Save(c Customer) error
}

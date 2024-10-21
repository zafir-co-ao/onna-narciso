package scheduling

type ServiceRepository interface {
	Get(id string) (Service, error)
	Save(s Service) error
}

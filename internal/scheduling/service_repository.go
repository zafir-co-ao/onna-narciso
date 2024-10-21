package scheduling

import "errors"

var ErrServiceNotFound = errors.New("Service not found")

type ServiceRepository interface {
	Get(id string) (Service, error)
	Save(s Service) error
}

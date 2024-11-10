package session

import (
	"errors"

	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

var ErrServiceNotFound = errors.New("service not found")

var EmptyServices = []Service{}

type ServiceAcl interface {
	FindByIDs(i []id.ID) ([]Service, error)
}

type ServiceAclFunc func(i []id.ID) ([]Service, error)

func (f ServiceAclFunc) FindByIDs(i []id.ID) ([]Service, error) {
	return f(i)
}

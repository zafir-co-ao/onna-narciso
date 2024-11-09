package session

import "github.com/zafir-co-ao/onna-narciso/internal/shared/id"

type ServiceAcl interface {
	FindByIDs(i []id.ID) ([]Service, error)
}

type ServiceAclFunc func(i []id.ID) ([]Service, error)

func (f ServiceAclFunc) FindByIDs(i []id.ID) ([]Service, error) {
	return f(i)
}

package stubs

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
)

func NewServicesServiceACL() scheduling.ServicesServiceACL {
	f := func(id nanoid.ID) (scheduling.Service, error) {
		for _, s := range testdata.Services {
			if s.ID == id {
				return s, nil
			}
		}
		return scheduling.Service{}, scheduling.ErrServiceNotFound
	}
	return scheduling.ServicesServiceACLFunc(f)
}

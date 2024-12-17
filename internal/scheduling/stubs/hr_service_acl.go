package stubs

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
)

func NewHRServiceACL() scheduling.HRServiceACL {
	f := func(id nanoid.ID) (scheduling.Professional, error) {
		for _, p := range testdata.Professionals {
			if p.ID == id {
				return p, nil
			}
		}

		return scheduling.Professional{}, scheduling.ErrProfessionalNotFound
	}

	return scheduling.HRServiceACLFunc(f)
}

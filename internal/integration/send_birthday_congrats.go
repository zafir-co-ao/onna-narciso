package integration

import (
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

func findBirthdayCustomers(finder crm.CustomerFinder, d date.Date) ([]crm.CustomerOutput, error) {

	day := d.String()[4:]

	customers, err := finder.FindAll()
	if err != nil {
		return nil, err
	}

	filtered := xslices.Filter(customers, func(c crm.CustomerOutput) bool {
		return c.BirthDate[4:] == day
	})

	return filtered, nil
}

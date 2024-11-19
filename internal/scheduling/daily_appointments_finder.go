package scheduling

import (
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

type DailyAppointmentsFinder interface {
	Find(theDate string) ([]AppointmentOutput, error)
}

type dailyAppointmentsFinderImpl struct {
	repo AppointmentRepository
}

func NewDailyAppointmentsFinder(repo AppointmentRepository) DailyAppointmentsFinder {
	return &dailyAppointmentsFinderImpl{repo: repo}
}

func (d *dailyAppointmentsFinderImpl) Find(theDate string) ([]AppointmentOutput, error) {
	a, err := d.repo.FindByDate(date.Date(theDate))

	if err != nil {
		return nil, err
	}

	return xslices.Map(a, toAppointmentOutput), nil
}

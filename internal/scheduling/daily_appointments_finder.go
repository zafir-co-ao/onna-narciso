package scheduling

import (
	"github.com/kindalus/godx/pkg/xslices"
	_date "github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

type DailyAppointmentsFinder interface {
	Find(date string) ([]AppointmentOutput, error)
}

type dailyAppointmentsFinderImpl struct {
	repo AppointmentRepository
}

func NewDailyAppointmentsGetter(repo AppointmentRepository) DailyAppointmentsFinder {
	return &dailyAppointmentsFinderImpl{repo: repo}
}

func (d *dailyAppointmentsFinderImpl) Find(date string) ([]AppointmentOutput, error) {
	a, err := d.repo.FindByDate(_date.Date(date))

	if err != nil {
		return nil, err
	}

	return xslices.Map(a, toAppointmentOutput), nil
}

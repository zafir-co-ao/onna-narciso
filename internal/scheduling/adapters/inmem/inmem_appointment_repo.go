package inmem

import (
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

type repo struct {
	data map[string]scheduling.Appointment
}

func NewAppointmentRepository() scheduling.AppointmentRepository {
	return &repo{
		data: make(map[string]scheduling.Appointment),
	}
}

func (r *repo) FindByID(id id.ID) (scheduling.Appointment, error) {

	if val, ok := r.data[id.Value()]; ok {
		return val, nil
	}

	return scheduling.EmptyAppointment, scheduling.ErrAppointmentNotFound
}

func (r *repo) Save(a scheduling.Appointment) error {
	r.data[a.ID.Value()] = a
	return nil
}

func (r *repo) FindByDate(d scheduling.Date) ([]scheduling.Appointment, error) {
	spec := scheduling.DateIsSpecificantion(d)

	var appointments []scheduling.Appointment
	for _, appointment := range r.data {
		if spec.IsSatisfiedBy(appointment) {
			appointments = append(appointments, appointment)
		}
	}

	return appointments, nil
}

func (r *repo) FindByWeekServiceAndProfessionals(date string, serviceID string, professionalsIDs []string) ([]scheduling.Appointment, error) {
	spec := shared.And(
		scheduling.WeekIsSpecificantion(date),
		scheduling.ServiceIsSpecificantion(serviceID),
		scheduling.ProfessionalsInSpecificantion(professionalsIDs),
	)

	var appointments []scheduling.Appointment
	for _, appointment := range r.data {
		if spec.IsSatisfiedBy(appointment) {
			appointments = append(appointments, appointment)
		}
	}

	return appointments, nil
}

func (r *repo) FindByDateAndStatus(date scheduling.Date, status scheduling.Status) ([]scheduling.Appointment, error) {
	spec := shared.And(
		scheduling.DateIsSpecificantion(date),
		scheduling.StatusIsSpecificantion(status),
	)

	var appointments []scheduling.Appointment
	for _, appointment := range r.data {
		if spec.IsSatisfiedBy(appointment) {
			appointments = append(appointments, appointment)
		}
	}

	return appointments, nil
}

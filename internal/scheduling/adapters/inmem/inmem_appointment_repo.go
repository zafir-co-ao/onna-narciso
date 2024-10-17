package inmem

import (
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

type repo struct {
	data map[string]scheduling.Appointment
}

func NewAppointmentRepository() scheduling.AppointmentRepository {
	return &repo{
		data: make(map[string]scheduling.Appointment),
	}
}

func (r *repo) Get(id string) (scheduling.Appointment, error) {

	if val, ok := r.data[id]; ok {
		return val, nil
	}

	return scheduling.Appointment{}, scheduling.ErrAppointmentNotFound
}

func (r *repo) Save(a scheduling.Appointment) error {
	r.data[a.ID] = a
	return nil
}

func (r *repo) FindByDate(d string) ([]scheduling.Appointment, error) {
	spec := scheduling.DateIsSpecificantion{Date: d}

	var appointments []scheduling.Appointment
	for _, appointment := range r.data {
		if spec.IsSatisfiedBy(appointment) {
			appointments = append(appointments, appointment)
		}
	}

	return appointments, nil
}

/*
func (r *repo) FindBySpecification(spec shared.Specification[scheduling.Appointment]) ([]scheduling.Appointment, error) {
	var appointments []scheduling.Appointment
	for _, appointment := range r.data {
		if spec.IsSatisfiedBy(appointment) {
			appointments = append(appointments, appointment)
		}
	}
	return appointments, nil
}
*/

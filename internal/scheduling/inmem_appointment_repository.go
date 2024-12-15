package scheduling

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

type inmemAppointmentRepositoryImpl struct {
	shared.BaseRepository[Appointment]
}

func NewAppointmentRepository(a ...Appointment) AppointmentRepository {
	return &inmemAppointmentRepositoryImpl{BaseRepository: shared.NewBaseRepository[Appointment](a...)}
}

func (r *inmemAppointmentRepositoryImpl) FindByID(id nanoid.ID) (Appointment, error) {
	if a, ok := r.Data[id]; ok {
		return a, nil
	}

	return EmptyAppointment, ErrAppointmentNotFound
}

func (r *inmemAppointmentRepositoryImpl) Save(a Appointment) error {
	r.Data[a.ID] = a
	return nil
}

func (r *inmemAppointmentRepositoryImpl) FindByDate(d date.Date) ([]Appointment, error) {
	spec := DateIsSpecification(d)

	var appointments []Appointment
	for _, appointment := range r.Data {
		if spec.IsSatisfiedBy(appointment) {
			appointments = append(appointments, appointment)
		}
	}

	return appointments, nil
}

func (r *inmemAppointmentRepositoryImpl) FindByWeekServiceAndProfessionals(date date.Date, serviceID nanoid.ID, professionalsIDs []nanoid.ID) ([]Appointment, error) {
	spec := shared.And(
		WeekIsSpecification(date),
		ServiceIsSpecification(serviceID),
		ProfessionalsIsSpecification(professionalsIDs),
		NotCanceledIsSpecification(),
	)

	var appointments []Appointment
	for _, a := range r.Data {
		if spec.IsSatisfiedBy(a) {
			appointments = append(appointments, a)
		}
	}

	return appointments, nil
}

func (r *inmemAppointmentRepositoryImpl) FindByDateStatusAndProfessional(date date.Date, status Status, id nanoid.ID) ([]Appointment, error) {
	spec := shared.And(
		DateIsSpecification(date),
		StatusIsSpecification(status),
		ProfessionalIsSpecification(id),
	)

	var appointments []Appointment
	for _, a := range r.Data {
		if spec.IsSatisfiedBy(a) {
			appointments = append(appointments, a)
		}
	}

	return appointments, nil
}

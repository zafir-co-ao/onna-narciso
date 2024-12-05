package inmem

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

type inmemAppointmentRepositoryImpl struct {
	shared.BaseRepository[scheduling.Appointment]
}

func NewAppointmentRepository(s ...scheduling.Appointment) scheduling.AppointmentRepository {
	r := &inmemAppointmentRepositoryImpl{
		BaseRepository: shared.NewBaseRepository[scheduling.Appointment](),
	}

	for _, a := range s {
		r.Save(a)
	}

	return r
}

func (r *inmemAppointmentRepositoryImpl) FindByID(id nanoid.ID) (scheduling.Appointment, error) {

	if val, ok := r.BaseRepository.Data[id]; ok {
		return val, nil
	}

	return scheduling.EmptyAppointment, scheduling.ErrAppointmentNotFound
}

func (r *inmemAppointmentRepositoryImpl) Save(a scheduling.Appointment) error {
	r.BaseRepository.Data[a.ID] = a
	return nil
}

func (r *inmemAppointmentRepositoryImpl) FindByDate(d date.Date) ([]scheduling.Appointment, error) {
	spec := scheduling.DateIsSpecification(d)

	var appointments []scheduling.Appointment
	for _, appointment := range r.BaseRepository.Data {
		if spec.IsSatisfiedBy(appointment) {
			appointments = append(appointments, appointment)
		}
	}

	return appointments, nil
}

func (r *inmemAppointmentRepositoryImpl) FindByWeekServiceAndProfessionals(date date.Date, serviceID nanoid.ID, professionalsIDs []nanoid.ID) ([]scheduling.Appointment, error) {
	spec := shared.And(
		scheduling.WeekIsSpecification(date),
		scheduling.ServiceIsSpecification(serviceID),
		scheduling.ProfessionalsInSpecification(professionalsIDs),
		scheduling.NotCanceledIsSpecification(),
	)

	var appointments []scheduling.Appointment
	for _, appointment := range r.BaseRepository.Data {
		if spec.IsSatisfiedBy(appointment) {
			appointments = append(appointments, appointment)
		}
	}

	return appointments, nil
}

func (r *inmemAppointmentRepositoryImpl) FindByDateStatusAndProfessional(date date.Date, status scheduling.Status, id nanoid.ID) ([]scheduling.Appointment, error) {
	spec := shared.And(
		scheduling.DateIsSpecification(date),
		scheduling.StatusIsSpecification(status),
		scheduling.ProfessionalIsSpecification(id),
	)

	var appointments []scheduling.Appointment
	for _, appointment := range r.BaseRepository.Data {
		if spec.IsSatisfiedBy(appointment) {
			appointments = append(appointments, appointment)
		}
	}

	return appointments, nil
}

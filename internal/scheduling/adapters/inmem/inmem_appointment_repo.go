package inmem

import (
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
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

func (r *inmemAppointmentRepositoryImpl) FindByID(id id.ID) (scheduling.Appointment, error) {

	if val, ok := r.Data[id]; ok {
		return val, nil
	}

	return scheduling.EmptyAppointment, scheduling.ErrAppointmentNotFound
}

func (r *inmemAppointmentRepositoryImpl) Save(a scheduling.Appointment) error {
	r.Data[a.ID] = a
	return nil
}

func (r *inmemAppointmentRepositoryImpl) FindByDate(d scheduling.Date) ([]scheduling.Appointment, error) {
	spec := scheduling.DateIsSpecificantion(d)

	var appointments []scheduling.Appointment
	for _, appointment := range r.Data {
		if spec.IsSatisfiedBy(appointment) {
			appointments = append(appointments, appointment)
		}
	}

	return appointments, nil
}

func (r *inmemAppointmentRepositoryImpl) FindByWeekServiceAndProfessionals(date string, serviceID string, professionalsIDs []string) ([]scheduling.Appointment, error) {
	spec := shared.And(
		scheduling.WeekIsSpecificantion(date),
		scheduling.ServiceIsSpecificantion(serviceID),
		scheduling.ProfessionalsInSpecificantion(professionalsIDs),
	)

	var appointments []scheduling.Appointment
	for _, appointment := range r.Data {
		if spec.IsSatisfiedBy(appointment) {
			appointments = append(appointments, appointment)
		}
	}

	return appointments, nil
}

func (r *inmemAppointmentRepositoryImpl) FindByDateAndStatus(date scheduling.Date, status scheduling.Status) ([]scheduling.Appointment, error) {
	spec := shared.And(
		scheduling.DateIsSpecificantion(date),
		scheduling.StatusIsSpecificantion(status),
	)

	var appointments []scheduling.Appointment
	for _, appointment := range r.BaseRepository.Data {
		if spec.IsSatisfiedBy(appointment) {
			appointments = append(appointments, appointment)
		}
	}

	return appointments, nil
}

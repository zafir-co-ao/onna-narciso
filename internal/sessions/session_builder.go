package sessions

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
)

type sessionBuilder Session

func NewSessionBuilder() *sessionBuilder {
	return &sessionBuilder{
		ID:          nanoid.New(),
		Status:      StatusCheckedIn,
		Date:        date.Today(),
		CheckinTime: hour.Now(),
		Services:    EmptyServices,
	}
}

func (s *sessionBuilder) WithID(id nanoid.ID) *sessionBuilder {
	s.ID = id
	return s
}

func (s *sessionBuilder) WithAppointmentID(appointmentID nanoid.ID) *sessionBuilder {
	s.AppointmentID = appointmentID
	return s
}

func (s *sessionBuilder) WithStatus(status Status) *sessionBuilder {
	s.Status = status
	return s
}

func (s *sessionBuilder) WithCheckinTime(checkinTime hour.Hour) *sessionBuilder {
	s.CheckinTime = checkinTime
	return s
}

func (s *sessionBuilder) WithCustomer(id nanoid.ID, name string) *sessionBuilder {
	s.CustomerID = id
	s.CustomerName = name
	return s
}

func (s *sessionBuilder) WithService(sid nanoid.ID, sname string, pid nanoid.ID, pname string) *sessionBuilder {

	ss := SessionService{
		ID:               sid,
		Name:             sname,
		ProfessionalID:   pid,
		ProfessionalName: pname,
	}

	s.Services = append(s.Services, ss)
	return s
}

func (s sessionBuilder) Build() Session {
	return Session(s)
}

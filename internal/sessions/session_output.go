package sessions

type SessionOutput struct {
	ID            string
	AppointmentID string
	Status        string
}

func toSessionOutput(s Session) SessionOutput {
	return SessionOutput{
		ID:            s.ID.String(),
		AppointmentID: s.AppointmentID.String(),
		Status:        string(s.Status),
	}
}

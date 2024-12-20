package sessions

import "github.com/kindalus/godx/pkg/xslices"

type SessionServiceOutput struct {
	ID       string
	Name     string
	Price    string
	Discount string
}

type SessionOutput struct {
	ID            string
	AppointmentID string
	Status        string
	Services      []SessionServiceOutput
}

func toSessionOutput(s Session) SessionOutput {
	return SessionOutput{
		ID:            s.ID.String(),
		AppointmentID: s.AppointmentID.String(),
		Status:        string(s.Status),
		Services: xslices.Map(s.Services, func(s SessionService) SessionServiceOutput {
			return SessionServiceOutput{
				ID:       s.ID.String(),
				Name:     s.Name,
				Price:    s.Price,
				Discount: s.Discount,
			}
		}),
	}
}

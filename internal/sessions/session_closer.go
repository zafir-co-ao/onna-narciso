package sessions

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

const EventSessionClosed = "EventSessionClosed"

type SessionCloserServiceInput struct {
	ServiceID string
	Discount  string
}

type SessionCloserInput struct {
	SessionID string
	Gift      string
	Services  []SessionCloserServiceInput
}

type SessionCloser interface {
	Close(i SessionCloserInput) error
}

type closerImpl struct {
	repo Repository
	sacl ServicesServiceACL
	bus  event.Bus
}

func NewSessionCloser(repo Repository, sacl ServicesServiceACL, bus event.Bus) SessionCloser {
	return &closerImpl{repo, sacl, bus}
}

func (u *closerImpl) Close(i SessionCloserInput) error {
	s, err := u.repo.FindByID(nanoid.ID(i.SessionID))
	if err != nil {
		return ErrSessionNotFound
	}

	svc, err := u.findServices(i.Services)
	if err != nil {
		return err
	}

	err = s.Close(svc, i.Gift)
	if err != nil {
		return err
	}

	err = u.repo.Save(s)
	if err != nil {
		return err
	}

	e := event.New(
		EventSessionClosed,
		event.WithHeader(event.HeaderAggregateID, s.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return nil
}

func (u *closerImpl) findServices(selectedServices []SessionCloserServiceInput) ([]SessionService, error) {
	if len(selectedServices) == 0 {
		return EmptyServices, nil
	}

	ids := xslices.Map(selectedServices, func(s SessionCloserServiceInput) nanoid.ID {
		return shared.StringToNanoid(s.ServiceID)
	})

	services, err := u.sacl.FindByIDs(ids)
	if err != nil {
		return []SessionService{}, err
	}

	for i, s := range selectedServices {
		if s.ServiceID == services[i].ID.String() {
			services[i].Discount = s.Discount
		}
	}

	return services, nil
}

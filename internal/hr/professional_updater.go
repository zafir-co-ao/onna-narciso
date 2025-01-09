package hr

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

var EventProfessionalUpdated = "EventProfessionalUpdated"

type ProfessionalUpdaterInput struct {
	ID          string
	Name        string
	ServicesIDs []string
}

type ProfessionalUpdater interface {
	Update(i ProfessionalUpdaterInput) error
}

type updaterImpl struct {
	repo Repository
	sacl ServicesServiceACL
	bus  event.Bus
}

func NewProfessionalUpdater(repo Repository, sacl ServicesServiceACL, bus event.Bus) ProfessionalUpdater {
	return &updaterImpl{repo, sacl, bus}
}

func (u *updaterImpl) Update(i ProfessionalUpdaterInput) error {
	_, err := u.repo.FindByID(nanoid.ID(i.ID))
	if err != nil {
		return err
	}

	ids := xslices.Map(i.ServicesIDs, func(id string) nanoid.ID {
		return shared.StringToNanoid(id)
	})

	s, err := u.sacl.FindServicesByIDs(ids)
	if err != nil {
		return err
	}

	name, err := name.New(i.Name)
	if err != nil {
		return err
	}

	p := NewProfessionalBuilder().
		WithID(nanoid.ID(i.ID)).
		WithName(name).
		WithServices(s).
		Build()

	err = u.repo.Save(p)
	if err != nil {
		return err
	}

	e := event.New(
		EventProfessionalUpdated,
		event.WithHeader(event.HeaderAggregateID, p.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return err
}

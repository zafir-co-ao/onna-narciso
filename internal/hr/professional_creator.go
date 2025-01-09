package hr

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type ProfessionalCreatorInput struct {
	Name        string
	ServicesIDs []string
}

type ProfessionalCreator interface {
	Create(i ProfessionalCreatorInput) (ProfessionalOutput, error)
}

func NewProfessionalCreator(repo Repository, sacl ServicesServiceACL, bus event.Bus) ProfessionalCreator {
	return &creatorImpl{repo, sacl, bus}
}

type creatorImpl struct {
	repo Repository
	sacl ServicesServiceACL
	bus  event.Bus
}

func (u *creatorImpl) Create(i ProfessionalCreatorInput) (ProfessionalOutput, error) {
	ids := xslices.Map(i.ServicesIDs, func(id string) nanoid.ID {
		return shared.StringToNanoid(id)
	})

	services, err := u.sacl.FindServicesByIDs(ids)
	if err != nil {
		return ProfessionalOutput{}, err
	}

	name, err := name.New(i.Name)
	if err != nil {
		return ProfessionalOutput{}, err
	}

	p := NewProfessionalBuilder().
		WithName(name).
		WithServices(services).
		Build()

	err = u.repo.Save(p)
	if err != nil {
		return ProfessionalOutput{}, err
	}

	e := event.New(
		EventProfessionalCreated,
		event.WithHeader(event.HeaderAggregateID, p.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return toProfessionalOutput(p), nil
}

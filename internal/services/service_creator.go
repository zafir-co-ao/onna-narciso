package services

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

const EventServiceCreated = "EventServiceCreated"

type ServiceCreatorInput struct {
	Name        string
	Description string
	Price       string
	Discount    string
	Duration    int
}

type ServiceCreator interface {
	Create(i ServiceCreatorInput) (ServiceOutput, error)
}

type serviceCreatorImpl struct {
	repo Repository
	bus  event.Bus
}

func NewServiceCreator(repo Repository, bus event.Bus) ServiceCreator {
	return &serviceCreatorImpl{repo, bus}
}

func (u *serviceCreatorImpl) Create(i ServiceCreatorInput) (ServiceOutput, error) {
	name, err := name.New(i.Name)
	if err != nil {
		return ServiceOutput{}, err
	}

	duration, err := duration.New(i.Duration)
	if err != nil {
		return ServiceOutput{}, err
	}

	price, err := NewPrice(i.Price)
	if err != nil {
		return ServiceOutput{}, err
	}

	discount, err := NewDiscount(i.Discount)
	if err != nil {
		return ServiceOutput{}, err
	}

	s := NewServiceBuilder().
		WithID(nanoid.New()).
		WithName(name).
		WithPrice(price).
		WithDuration(duration).
		WithDiscount(discount).
		WithDescription(Description(i.Description)).
		Build()

	err = u.repo.Save(s)
	if err != nil {
		return ServiceOutput{}, err
	}

	e := event.New(
		EventServiceCreated,
		event.WithHeader(event.HeaderAggregateID, s.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return toServiceOutput(s), nil
}

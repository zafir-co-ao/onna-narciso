package services

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/services/price"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

const EventServiceCreated = "EventServiceCreated"

type ServiceCreatorInput struct {
	Name        string
	Description string
	Price       string
	Duration    int
}

type ServiceOutput struct {
	ID          string
	Name        string
	Description string
	Price       string
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
	_name, err := name.New(i.Name)
	if err != nil {
		return ServiceOutput{}, err
	}

	_duration, err := duration.New(i.Duration)
	if err != nil {
		return ServiceOutput{}, err
	}

	_price, err := price.New(i.Price)
	if err != nil {
		return ServiceOutput{}, err
	}

	s := NewService(
		_name,
		_duration,
		_price,
		Description(i.Description),
	)

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

func toServiceOutput(s Service) ServiceOutput {
	return ServiceOutput{
		ID:          s.ID.String(),
		Name:        s.Name.String(),
		Duration:    s.Duration.Value(),
		Description: string(s.Description),
		Price:       string(s.Price),
	}
}

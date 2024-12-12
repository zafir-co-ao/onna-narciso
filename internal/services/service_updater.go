package services

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

var EventPriceUpdated = "EventPriceUpdated"

type ServiceUpdaterInput struct {
	ID          string
	Name        string
	Description string
	Price       string
	Duration    int
}

type ServiceUpdater interface {
	Update(i ServiceUpdaterInput) error
}

type updaterImpl struct {
	repo Repository
	bus  event.Bus
}

func NewServiceUpdater(repo Repository, bus event.Bus) ServiceUpdater {
	return &updaterImpl{repo, bus}
}

func (u *updaterImpl) Update(i ServiceUpdaterInput) error {
	s, err := u.repo.FindByID(nanoid.ID(i.ID))
	if err != nil {
		return err
	}

	_name, err := name.New(i.Name)
	if err != nil {
		return err
	}

	_price, err := NewPrice(i.Price)
	if err != nil {
		return err
	}

	_duration, err := duration.New(i.Duration)
	if err != nil {
		return err
	}

	oldPrice := s.Price

	s = NewServiceBuilder().
		WithID(nanoid.ID(i.ID)).
		WithName(_name).
		WithPrice(_price).
		WithDuration(_duration).
		WithDescription(Description(i.Description)).
		Build()

	err = u.repo.Save(s)
	if err != nil {
		return err
	}

	if s.IsSamePrice(oldPrice) {
		return nil
	}

	e := event.New(
		EventPriceUpdated,
		event.WithHeader(event.HeaderAggregateID, s.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return nil
}

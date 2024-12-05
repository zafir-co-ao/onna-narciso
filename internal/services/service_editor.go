package services

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/services/price"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

const EventServiceEdited = "EventServiceEdited"

type ServiceEditorInput struct {
	ID          string
	Name        string
	Description string
	Price       string
	Duration    int
}

type ServiceEditor interface {
	Edit(i ServiceEditorInput) error
}

type ServiceEditorImpl struct {
	repo Repository
	bus  event.Bus
}

func NewServiceEditor(repo Repository, bus event.Bus) ServiceEditor {
	return &ServiceEditorImpl{repo: repo, bus: bus}
}

func (u *ServiceEditorImpl) Edit(i ServiceEditorInput) error {

	_, err := u.repo.FindByID(nanoid.ID(i.ID))
	if err != nil {
		return err
	}

	_name, err := name.New(i.Name)
	if err != nil {
		return err
	}

	_price, err := price.New(i.Price)
	if err != nil {
		return err
	}

	_duration, err := duration.New(i.Duration)
	if err != nil {
		return err
	}

	s := NewServiceBuilder().WithID(nanoid.ID(i.ID)).
		WithName(_name).
		WithPrice(_price).
		WithDuration(_duration).
		WithDescription(Description(i.Description)).
		Build()

	err = u.repo.Save(s)
	if err != nil {
		return err
	}

	e := event.New(
		EventServiceEdited,
		event.WithHeader(event.HeaderAggregateID, s.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return nil
}

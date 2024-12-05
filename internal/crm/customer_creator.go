package crm

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

const EventCustomerCreated = "EventCustomerCreated"

type CustomerCreatornput struct {
	Name        string
	Nif         string
	BirthDate   string
	Email       string
	PhoneNumber string
}

type CustomerCreator interface {
	Create(i CustomerCreatornput) (CustomerOutput, error)
}

type customerCreatorImpl struct {
	repo Repository
	bus  event.Bus
}

func NewCustomerCreator(repo Repository, bus event.Bus) CustomerCreator {
	return &customerCreatorImpl{repo, bus}
}

func (u *customerCreatorImpl) Create(i CustomerCreatornput) (CustomerOutput, error) {
	_, err := u.repo.FindByNif(Nif(i.Nif))
	if err != nil {
		return CustomerOutput{}, err
	}

	c := NewCustomerBuilder().
		WithName(name.Name(i.Name)).
		WithNif(Nif(i.Nif)).
		WithBirthDate(date.Date(i.BirthDate)).
		WithEmail(i.Email).
		WithPhoneNumber(i.PhoneNumber).
		Build()

	err = u.repo.Save(c)
	if err != nil {
		return CustomerOutput{}, err
	}

	e := event.New(
		EventCustomerCreated,
		event.WithHeader(event.HeaderAggregateID, c.ID.String()),
		event.WithPayload(c),
	)

	u.bus.Publish(e)

	return toCustomerOutput(c), nil
}

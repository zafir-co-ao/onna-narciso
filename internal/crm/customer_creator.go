package crm

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

const EventCustomerCreated = "EventCustomerCreated"

type CustomerCreatorInput struct {
	Name        string
	Nif         string
	BirthDate   string
	Email       string
	PhoneNumber string
}

type CustomerCreator interface {
	Create(i CustomerCreatorInput) (CustomerOutput, error)
}

type customerCreatorImpl struct {
	repo Repository
	bus  event.Bus
}

func NewCustomerCreator(repo Repository, bus event.Bus) CustomerCreator {
	return &customerCreatorImpl{repo, bus}
}

func (u *customerCreatorImpl) Create(i CustomerCreatorInput) (CustomerOutput, error) {
	nif, err := NewNif(i.Nif)
	if err != nil {
		return CustomerOutput{}, err
	}

	_, err = u.repo.FindByNif(nif)
	if err != nil {
		return CustomerOutput{}, err
	}

	n, err := name.New(i.Name)
	if err != nil {
		return CustomerOutput{}, err
	}

	if !date.IsValidFormat(i.BirthDate) {
		return CustomerOutput{}, date.ErrInvalidFormat
	}

	email, err := NewEmail(i.Email)
	if err != nil {
		return CustomerOutput{}, err
	}

	p, err := NewPhoneNumber(i.PhoneNumber)
	if err != nil {
		return CustomerOutput{}, err
	}

	c := NewCustomerBuilder().
		WithName(n).
		WithNif(nif).
		WithBirthDate(date.Date(i.BirthDate)).
		WithEmail(email).
		WithPhoneNumber(p).
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

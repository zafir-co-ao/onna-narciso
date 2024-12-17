package crm

import (
	"github.com/kindalus/godx/pkg/event"
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
	name, err := name.New(i.Name)
	if err != nil {
		return CustomerOutput{}, err
	}

	nif, err := NewNif(i.Nif)
	if err != nil {
		return CustomerOutput{}, err
	}

	bdate, err := getBirthDate(i.BirthDate)
	if err != nil {
		return CustomerOutput{}, err
	}

	if !isAllowedAge(bdate) {
		return CustomerOutput{}, ErrAgeNotAllowed
	}

	email, err := NewEmail(i.Email)
	if err != nil {
		return CustomerOutput{}, err
	}

	phone, err := NewPhoneNumber(i.PhoneNumber)
	if err != nil {
		return CustomerOutput{}, err
	}

	customers, err := u.repo.FindAll()

	if err != nil {
		return CustomerOutput{}, err
	}

	if checkUsedNif(customers, nif) {
		return CustomerOutput{}, ErrNifAlreadyUsed
	}

	if checkUsedEmail(customers, email) {
		return CustomerOutput{}, ErrEmailAlreadyUsed
	}

	if checkUsedPhoneNumber(customers, phone) {
		return CustomerOutput{}, ErrPhoneNumberAlreadyUsed
	}

	c := NewCustomerBuilder().
		WithName(name).
		WithNif(nif).
		WithBirthDate(bdate).
		WithEmail(email).
		WithPhoneNumber(phone).
		Build()

	err = u.repo.Save(c)
	if err != nil {
		return CustomerOutput{}, err
	}

	e := event.New(
		EventCustomerCreated,
		event.WithHeader(event.HeaderAggregateID, c.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return toCustomerOutput(c), nil
}

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

	b, err := date.New(i.BirthDate)
	if err != nil {
		return CustomerOutput{}, err
	}

	if !isAllowedAge(b) {
		return CustomerOutput{}, ErrAgeNotAllowed
	}

	email, err := NewEmail(i.Email)
	if err != nil {
		return CustomerOutput{}, err
	}

	if u.isUsedEmail(email) {
		return CustomerOutput{}, ErrEmailAlreadyUsed
	}

	p, err := NewPhoneNumber(i.PhoneNumber)
	if err != nil {
		return CustomerOutput{}, err
	}

	if u.isUsedPhoneNumber(p) {
		return CustomerOutput{}, ErrPhoneNumberAlreadyUsed
	}

	c := NewCustomerBuilder().
		WithName(n).
		WithNif(nif).
		WithBirthDate(b).
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

func (u *customerCreatorImpl) isUsedPhoneNumber(number PhoneNumber) bool {
	if len(number) == 0 {
		return false
	}

	_, err := u.repo.FindByPhoneNumber(number)
	return err == nil
}

func (u *customerCreatorImpl) isUsedEmail(email Email) bool {
	if len(email) == 0 {
		return false
	}

	_, err := u.repo.FindByEmail(email)
	return err == nil
}

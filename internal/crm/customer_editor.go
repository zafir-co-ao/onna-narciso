package crm

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

const EventCustomerUpdated = "EventCustomerUpdated"

type CustomerEditorInput struct {
	ID          nanoid.ID
	Name        string
	Nif         string
	BirthDate   string
	Email       string
	PhoneNumber string
}

type CustomerEditor interface {
	Edit(i CustomerEditorInput) error
}

type customerEditorImpl struct {
	repo Repository
	bus  event.Bus
}

func NewCustomerEditor(repo Repository, bus event.Bus) CustomerEditor {
	return &customerEditorImpl{repo: repo, bus: bus}
}

func (u *customerEditorImpl) Edit(i CustomerEditorInput) error {
	c, _ := u.repo.FindByID(i.ID)

	if u.isUsedNif(c, i) {
		return ErrNifAlreadyUsed
	}

	name, err := name.New(i.Name)
	if err != nil {
		return err
	}

	nif, err := NewNif(i.Nif)
	if err != nil {
		return err
	}

	email, err := NewEmail(i.Email)
	if err != nil {
		return err
	}

	p, err := NewPhoneNumber(i.PhoneNumber)
	if err != nil {
		return err
	}

	if !date.IsValidFormat(i.BirthDate) {
		return date.ErrInvalidFormat
	}

	c = NewCustomerBuilder().
		WithID(i.ID).
		WithName(name).
		WithNif(nif).
		WithBirthDate(date.Date(i.BirthDate)).
		WithEmail(email).
		WithPhoneNumber(p).
		Build()

	u.repo.Save(c)

	e := event.New(
		EventCustomerUpdated,
		event.WithHeader(event.HeaderAggregateID, c.ID.String()),
		event.WithPayload(c),
	)

	u.bus.Publish(e)

	return nil
}

func (u *customerEditorImpl) isUsedNif(c Customer, i CustomerEditorInput) bool {
	if c.IsSameNif(Nif(i.Nif)) {
		return false
	}

	_, err := u.repo.FindByNif(Nif(i.Nif))
	return err != nil
}

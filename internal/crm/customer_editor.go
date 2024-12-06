package crm

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/email"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/nif"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/phone"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

const EventCustomerUpdated = "EventCustomerUpdated"

type CustomerEditorInput struct {
	ID          nanoid.ID
	Name        string
	NIF         string
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

	name, _ := name.New(i.Name)

	nif, _ := nif.New(i.NIF)

	email, err := email.New(i.Email)
	if err != nil {
		return err
	}

	p, _ := phone.New(i.PhoneNumber)

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
	if c.IsSameNif(nif.Nif(i.NIF)) {
		return false
	}

	_, err := u.repo.FindByNif(nif.Nif(i.NIF))
	return err != nil
}

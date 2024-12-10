package crm

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

const EventCustomerUpdated = "EventCustomerUpdated"

type CustomerEditorInput struct {
	ID          string
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
	c, err := u.repo.FindByID(nanoid.ID(i.ID))
	if err != nil {
		return err
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

	if u.isUsedNif(c, i.Nif) {
		return ErrNifAlreadyUsed
	}

	if u.isUsedEmail(c, i.Email) {
		return ErrEmailAlreadyUsed
	}

	if u.isUsedPhoneNumber(c, i.PhoneNumber) {
		return ErrPhoneNumberAlreadyUsed
	}

	b, _ := date.New(i.BirthDate)

	if !isAllowedAge(b) {
		return ErrAgeNotAllowed
	}

	c = NewCustomerBuilder().
		WithID(nanoid.ID(i.ID)).
		WithName(name).
		WithNif(nif).
		WithBirthDate(date.Date(i.BirthDate)).
		WithEmail(email).
		WithPhoneNumber(p).
		Build()

	err = u.repo.Save(c)
	if err != nil {
		return err
	}

	e := event.New(
		EventCustomerUpdated,
		event.WithHeader(event.HeaderAggregateID, c.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return nil
}

func (u *customerEditorImpl) isUsedNif(c Customer, nif string) bool {
	if c.IsSameNif(Nif(nif)) {
		return false
	}

	_, err := u.repo.FindByNif(Nif(nif))
	return err != nil
}

func (u *customerEditorImpl) isUsedEmail(c Customer, email string) bool {
	if len(email) == 0 {
		return false
	}

	if c.IsSameEmail(Email(email)) {
		return false
	}

	_, err := u.repo.FindByEmail(Email(email))
	return err == nil
}

func (u *customerEditorImpl) isUsedPhoneNumber(c Customer, phoneNumber string) bool {
	if len(phoneNumber) == 0 {
		return false
	}

	if c.IsSamePhoneNumber(PhoneNumber(phoneNumber)) {
		return false
	}

	_, err := u.repo.FindByPhoneNumber(PhoneNumber(phoneNumber))
	return err == nil
}

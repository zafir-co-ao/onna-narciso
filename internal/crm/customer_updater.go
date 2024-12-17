package crm

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

const EventCustomerUpdated = "EventCustomerUpdated"

type CustomerUpdaterInput struct {
	ID          string
	Name        string
	Nif         string
	BirthDate   string
	Email       string
	PhoneNumber string
}

type CustomerUpdater interface {
	Update(i CustomerUpdaterInput) error
}

type updaterImpl struct {
	repo Repository
	bus  event.Bus
}

func NewCustomerUpdater(repo Repository, bus event.Bus) CustomerUpdater {
	return &updaterImpl{repo, bus}
}

func (u *updaterImpl) Update(i CustomerUpdaterInput) error {
	c, err := u.repo.FindByID(nanoid.ID(i.ID))
	if err != nil {
		return err
	}

	name, err := name.New(i.Name)
	if err != nil {
		return err
	}

	email, err := NewEmail(i.Email)
	if err != nil {
		return err
	}

	bdate, err := getBirthDate(i.BirthDate)
	if err != nil {
		return ErrAgeNotAllowed
	}

	if !isAllowedAge(bdate) {
		return ErrAgeNotAllowed
	}

	customers, err := u.repo.FindAll()
	if err != nil {
		return err
	}

	customers = xslices.Filter(customers, func(customer Customer) bool {
		return customer.ID != c.ID
	})

	if checkUsedPhoneNumber(customers, PhoneNumber(i.PhoneNumber)) {
		return ErrPhoneNumberAlreadyUsed
	}

	if checkUsedEmail(customers, email) {
		return ErrEmailAlreadyUsed
	}

	if checkUsedNif(customers, Nif(i.Nif)) {
		return ErrNifAlreadyUsed
	}

	c = NewCustomerBuilder().
		WithID(nanoid.ID(i.ID)).
		WithName(name).
		WithNif(Nif(i.Nif)).
		WithBirthDate(bdate).
		WithEmail(email).
		WithPhoneNumber(PhoneNumber(i.PhoneNumber)).
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

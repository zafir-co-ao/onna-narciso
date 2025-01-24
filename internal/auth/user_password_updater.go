package auth

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
)

const EventUserPasswordUpdated = "EventUserPasswordUpdated"

type UserPasswordUpdater interface {
	Update(i UserPasswordUpdaterInput) error
}

type UserPasswordUpdaterInput struct {
	UserID   string
	Password string
}

func NewUserPasswordUpdater(r Repository, bus event.Bus) UserPasswordUpdater {
	return &userPasswordUpdaterImpl{r, bus}
}

type userPasswordUpdaterImpl struct {
	repo Repository
	bus  event.Bus
}

func (u *userPasswordUpdaterImpl) Update(i UserPasswordUpdaterInput) error {
	user, err := u.repo.FindByID(nanoid.ID(i.UserID))
	if err != nil {
		return err
	}

	password, err := NewPassword(i.Password)
	if err != nil {
		return err
	}

	user.SetPassword(password)

	err = u.repo.Save(user)
	if err != nil {
		return err
	}

	e := event.New(EventUserPasswordUpdated,
		event.WithHeader(event.HeaderAggregateID, string(user.ID)),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return nil
}

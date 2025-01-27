package auth

import (
	"errors"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
)

const EventUserPasswordUpdated = "EventUserPasswordUpdated"

var (
	ErrInvalidOldPassword          = errors.New("invalid old password")
	ErrInvalidConfirmationPassword = errors.New("invalid confirmation password")
)

type UserPasswordUpdater interface {
	Update(i UserPasswordUpdaterInput) error
}

type UserPasswordUpdaterInput struct {
	UserID               string
	OldPassword          string
	NewPassword          string
	ConfirmationPassword string
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

	if !user.VerifyPassword(i.OldPassword) {
		return ErrInvalidOldPassword
	}

	if !user.IsSamePassword(i.NewPassword, i.ConfirmationPassword) {
		return ErrInvalidConfirmationPassword
	}

	password, err := NewPassword(i.NewPassword)
	if err != nil {
		return err
	}

	user.UpdatePassword(password)

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

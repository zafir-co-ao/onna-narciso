package auth

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
)

type UserUpdaterInput struct {
	ManagerID   string
	UserID      string
	Username    string
	Email       string
	PhoneNumber string
	Role        string
}

const EventUserUpdated = "EventUserUpdated"

type UserUpdater interface {
	Update(i UserUpdaterInput) error
}

func NewUserUpdater(r Repository, bus event.Bus) UserUpdater {
	return &userUpdaterImpl{r, bus}
}

type userUpdaterImpl struct {
	repo Repository
	bus  event.Bus
}

func (u *userUpdaterImpl) Update(i UserUpdaterInput) error {
	manager, err := u.repo.FindByID(nanoid.ID(i.ManagerID))
	if err != nil {
		return err
	}

	if !manager.IsManager() {
		return ErrUserNotAllowed
	}

	_, err = u.repo.FindByID(nanoid.ID(i.UserID))
	if err != nil {
		return err
	}

	users, err := u.repo.FindAll()
	if err != nil {
		return err
	}

	users = xslices.Filter(users, func(u User) bool {
		return u.ID != nanoid.ID(i.UserID)
	})

	username, err := NewUsername(i.Username)
	if err != nil {
		return err
	}

	if !IsAvailableUsername(users, username) {
		return ErrOnlyUniqueUsername
	}

	email, err := NewEmail(i.Email)
	if err != nil {
		return err
	}

	if !IsAvailableEmail(users, email) {
		return ErrOnlyUniqueEmail
	}

	phoneNumber, err := NewPhoneNumber(i.PhoneNumber)
	if err != nil {
		return err
	}

	if !IsAvailablePhoneNumber(users, phoneNumber) {
		return ErrOnlyUniquePhoneNumber
	}

	role, err := NewRole(i.Role)
	if err != nil {
		return err
	}

	user := NewUserBuilder().
		WithID(nanoid.ID(i.UserID)).
		WithUserName(username).
		WithEmail(email).
		WithPhoneNumber(phoneNumber).
		WithRole(role).
		Build()

	if err := u.repo.Save(user); err != nil {
		return err
	}

	e := event.New(EventUserUpdated,
		event.WithHeader(event.HeaderAggregateID, user.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return nil
}

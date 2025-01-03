package auth

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
)

const EventUserCreated = "EventUserCreated"

type UserCreatorInput struct {
	UserID   string
	Username string
	Password string
	Role     string
}

type UserCreator interface {
	Create(i UserCreatorInput) (UserOutput, error)
}

type creatorImpl struct {
	repo Repository
	bus  event.Bus
}

func NewUserCreator(repo Repository, bus event.Bus) UserCreator {
	return &creatorImpl{repo, bus}
}

func (u *creatorImpl) Create(i UserCreatorInput) (UserOutput, error) {
	au, err := u.repo.FindByID(nanoid.ID(i.UserID))
	if err != nil {
		return UserOutput{}, err
	}

	if !au.IsManager() {
		return UserOutput{}, ErrUserNotAllowed
	}

	username, err := NewUsername(i.Username)
	if err != nil {
		return UserOutput{}, err
	}

	users, err := u.repo.FindAll()
	if err != nil {
		return UserOutput{}, err
	}

	if !u.isAvailableUsername(users, username) {
		return UserOutput{}, ErrOnlyUniqueUsername
	}

	password, err := NewPassword(i.Password)
	if err != nil {
		return UserOutput{}, err
	}

	role, err := NewRole(i.Role)
	if err != nil {
		return UserOutput{}, err
	}

	user := NewUser(username, password, role)

	err = u.repo.Save(user)
	if err != nil {
		return UserOutput{}, err
	}

	e := event.New(
		EventUserCreated,
		event.WithHeader(event.HeaderAggregateID, user.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return toUserOutput(user), nil
}

func (u *creatorImpl) isAvailableUsername(users []User, username Username) bool {
	return xslices.All(users, func(u User) bool {
		return u.Username != username
	})
}

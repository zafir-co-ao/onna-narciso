package auth

import (
	"fmt"
	"log/slog"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/notifications"
)

type UserPasswordResetter interface {
	Reset(i UserPasswordResetterInput) error
}

type UserPasswordResetterInput struct {
	Email string
}

const (
	PasswordLength                = 12
	EventUserPasswordReset        = "EventUserPasswordReset"
	ErrUserPasswordNotSentToEmail = "Erro ao enviar nova palavra para %s: %v"
)

func NewUserPasswordResetter(r Repository, bus event.Bus, n notifications.Notifier) UserPasswordResetter {
	return &userPasswordResetterImpl{repo: r, bus: bus, notifier: n}
}

type userPasswordResetterImpl struct {
	repo     Repository
	bus      event.Bus
	notifier notifications.Notifier
}

func (u *userPasswordResetterImpl) Reset(i UserPasswordResetterInput) error {
	email, err := NewEmail(i.Email)
	if err != nil {
		return err
	}

	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return err
	}

	p, err := GeneratePassword(PasswordLength)
	if err != nil {
		return err
	}

	user.ResetPassword(p)

	if err := u.repo.Save(user); err != nil {
		return err
	}

	if err = u.SendPasswordToEmail(user, p); err != nil {
		slog.Error(ErrUserPasswordNotSentToEmail, user.Email.String(), err)
	}

	e := event.New(EventUserPasswordReset,
		event.WithHeader(event.HeaderAggregateID, user.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return nil
}

func (u *userPasswordResetterImpl) SendPasswordToEmail(user User, p string) error {
	c := notifications.Contact{
		Name:  user.Username.String(),
		Email: user.Email.String(),
	}

	msg := notifications.Message{
		Subject: "Nova palavra-passe",
		Body:    fmt.Sprintf("A sua nova palavra-passe é: %s", p),
	}

	err := u.notifier.Notify(c, msg)
	if err != nil {
		return err
	}

	return nil
}

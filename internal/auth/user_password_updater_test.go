package auth_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)

func TestTestUserPasswordUpdater(t *testing.T) {
	user := auth.User{ID: "1", Password: auth.MustNewPassword("Konnely")}

	t.Run("should_retrieve_user_in_repository", func(t *testing.T) {
		repo := auth.NewInmemRepository(user)
		bus := event.NewEventBus()
		u := auth.NewUserPasswordUpdater(repo, bus)
		i := auth.UserPasswordUpdaterInput{
			UserID:               "1",
			NewPassword:          "gustin",
			ConfirmationPassword: "gustin",
			OldPassword:          "Konnely",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("should find user in repository, got %v", err)
		}
	})

	t.Run("should_update_password", func(t *testing.T) {
		repo := auth.NewInmemRepository(user)
		bus := event.NewEventBus()
		u := auth.NewUserPasswordUpdater(repo, bus)
		i := auth.UserPasswordUpdaterInput{
			UserID:               "1",
			NewPassword:          "123456",
			ConfirmationPassword: "123456",
			OldPassword:          "Konnely",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		user, err := repo.FindByID(nanoid.ID(i.UserID))

		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("should find user in repository, got %v", err)
		}

		if !user.VerifyPassword(i.NewPassword) {
			t.Error("Expected no error, got password not updated")
		}
	})

	t.Run("should_publish_domain_event_when_user_password_updated", func(t *testing.T) {
		repo := auth.NewInmemRepository(user)
		bus := event.NewEventBus()
		u := auth.NewUserPasswordUpdater(repo, bus)
		isPublished := false

		var h event.HandlerFunc = func(e event.Event) {
			if e.Name() == auth.EventUserPasswordUpdated {
				isPublished = true
			}
		}

		bus.Subscribe(auth.EventUserPasswordUpdated, h)

		i := auth.UserPasswordUpdaterInput{
			UserID:               "1",
			NewPassword:          "123456",
			ConfirmationPassword: "123456",
			OldPassword:          "Konnely",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("The %s must be published", auth.EventUserPasswordUpdated)
		}
	})

	t.Run("should_return_error_when_password_not_provided", func(t *testing.T) {
		repo := auth.NewInmemRepository(user)
		bus := event.NewEventBus()
		u := auth.NewUserPasswordUpdater(repo, bus)
		i := auth.UserPasswordUpdaterInput{
			UserID:               "1",
			NewPassword:          "",
			ConfirmationPassword: "",
			OldPassword:          "Konnely",
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		if !errors.Is(err, auth.ErrEmptyPassword) {
			t.Errorf("The error must be, got %v", err)
		}
	})

	t.Run("should_return_error_when_old_password_provided_is_incorrect", func(t *testing.T) {
		repo := auth.NewInmemRepository(user)
		bus := event.NewEventBus()
		u := auth.NewUserPasswordUpdater(repo, bus)
		i := auth.UserPasswordUpdaterInput{
			UserID:               "1",
			NewPassword:          "mypass",
			ConfirmationPassword: "mypass",
			OldPassword:          "Konne",
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		if !errors.Is(err, auth.ErrInvalidOldPassword) {
			t.Errorf("The error must be, got %v", err)
		}
	})

	t.Run("should_return_error_if_new_password_is_not_the_same_as_confirmation_password", func(t *testing.T) {
		repo := auth.NewInmemRepository(user)
		bus := event.NewEventBus()
		u := auth.NewUserPasswordUpdater(repo, bus)
		i := auth.UserPasswordUpdaterInput{
			UserID:               "1",
			OldPassword:          "Konnely",
			NewPassword:          "mypass",
			ConfirmationPassword: "mypas",
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		if !errors.Is(err, auth.ErrInvalidConfirmationPassword) {
			t.Errorf("The error must be, got %v", err)
		}
	})
}

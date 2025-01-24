package auth_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)

func TestTestUserPasswordUpdater(t *testing.T) {
	user := auth.User{ID: "1", Password: "Konnely"}
	repo := auth.NewInmemRepository(user)
	bus := event.NewEventBus()
	u := auth.NewUserPasswordUpdater(repo, bus)

	t.Run("should_update_user_password", func(t *testing.T) {
		i := auth.UserPasswordUpdaterInput{
			UserID: "1",
			Password: "kyle",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("should_retrieve_user_in_repository", func(t *testing.T) {
		i := auth.UserPasswordUpdaterInput{
			UserID: "1",
			Password: "gustin",
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
		i := auth.UserPasswordUpdaterInput{
			UserID:   "1",
			Password: "123456",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		user, err := repo.FindByID(nanoid.ID(i.UserID))

		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("should find user in repository, got %v", err)
		}

		if !user.VerifyPassword(i.Password) {
			t.Error("Expected no error, got password not updated")
		}
	})

	t.Run("should_publish_domain_event_when_user_password_updated", func(t *testing.T) {
		isPublished := false

		var h event.HandlerFunc = func(e event.Event) {
			if e.Name() == auth.EventUserPasswordUpdated {
				isPublished = true
			}
		}

		bus.Subscribe(auth.EventUserPasswordUpdated, h)

		i := auth.UserPasswordUpdaterInput{
			UserID:   "1",
			Password: "123456",
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
		i := auth.UserPasswordUpdaterInput{
			UserID:   "1",
			Password: "",
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		if !errors.Is(err, auth.ErrEmptyPassword) {
			t.Errorf("The error must be, got %v", err)
		}
	})
}

package auth_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/internal/auth/stubs"
)

func TestUserPasswordResetter(t *testing.T) {
	user := auth.User{Email: "kate@gmail.com", Password: auth.MustNewPassword("kate1234")}
	repo := auth.NewInmemRepository(user)
	nacl := stubs.NewNotificationsACL()
	bus := event.NewEventBus()

	u := auth.NewUserPasswordResetter(repo, bus, nacl)
	i := auth.UserPasswordResetterInput{Email: "kate@gmail.com"}

	t.Run("should_retrieve_user_in_repository", func(t *testing.T) {
		err := u.Reset(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if errors.Is(err, auth.ErrUserNotFound) {
			t.Error("should find user in repository")
		}
	})

	t.Run("should_reset_password", func(t *testing.T) {
		err := u.Reset(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		_user, err := repo.FindByEmail(auth.Email(i.Email))

		if errors.Is(err, auth.ErrUserNotFound) {
			t.Error("should find user in repository")
		}

		if _user.Password == user.Password {
			t.Errorf("should update password, got %v", _user.Password.String())
		}
	})

	t.Run("should_send_new_password_to_email", func(t *testing.T) {
		err := u.Reset(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if nacl.Contact.Email != user.Email.String() {
			t.Errorf("Expected email to be sent to %s, got %v", user.Email.String(), nacl.Contact.Email)
		}

		expectedPrefix := "A sua nova palavra-passe é:"
		if !strings.HasPrefix(nacl.Message.Body, expectedPrefix) {
			t.Errorf("Expected email body to contain new password prefix, got %s", nacl.Message.Body)
		}
	})

	t.Run("should_publish_event_when_user_password_is_reset", func(t *testing.T) {
		isPublished := false

		var h event.HandlerFunc = func(e event.Event) {
			if e.Name() == auth.EventUserPasswordReset {
				isPublished = true
			}
		}
		bus.Subscribe(auth.EventUserPasswordReset, h)

		err := u.Reset(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("The %s must be published", auth.EventUserPasswordReset)
		}
	})

	t.Run("should_return_error_if_email_is_empty", func(t *testing.T) {
		i := auth.UserPasswordResetterInput{Email: ""}

		err := u.Reset(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !errors.Is(err, auth.ErrEmptyEmail) {
			t.Errorf("The error must be %v, got %v", auth.ErrEmptyEmail, err)
		}
	})

	t.Run("should_return_error_if_email_is_in_invalid_format", func(t *testing.T) {
		i := auth.UserPasswordResetterInput{Email: "kate@123"}

		err := u.Reset(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !errors.Is(err, auth.ErrInvalidEmailFormat) {
			t.Errorf("The error must be %v, got %v", auth.ErrInvalidEmailFormat, err)
		}
	})

	t.Run("should_return_error_when_user_not_exists_in_repository", func(t *testing.T) {
		i := auth.UserPasswordResetterInput{Email: "kate1234@gmail.com"}

		err := u.Reset(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !errors.Is(err, auth.ErrUserNotFound) {
			t.Error("should find user in repository")
		}
	})
}

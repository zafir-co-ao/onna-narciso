package auth_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)

func TestUserAuthenticator(t *testing.T) {
	password, _ := auth.NewPassword("1234")
	user := auth.User{
		Username: "Jonathan",
		Password: password,
	}

	repo := auth.NewInmemRepository(user)
	u := auth.NewUserAuthenticator(repo)

	t.Run("should_authenticate_user", func(t *testing.T) {
		i := auth.UserAuthenticatorInput{
			Username: "Jonathan",
			Password: "1234",
		}

		_, err := u.Authenticate(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_authenticate_with_credendentials", func(t *testing.T) {
		i := auth.UserAuthenticatorInput{Username: "Jonathan", Password: "1234"}

		o, err := u.Authenticate(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.Username == "" {
			t.Errorf("Expected an user name, got %v", o.Username)
		}

		if o.Password == "" {
			t.Errorf("Expected a password, got %v", o.Password)
		}
	})

	t.Run("should_verify_credentials_of_user", func(t *testing.T) {
		i := auth.UserAuthenticatorInput{
			Username: "Jonathan",
			Password: "1234",
		}

		_, err := u.Authenticate(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		user, _ := repo.FindByUsername(auth.Username(i.Username))

		if user.Username.String() != i.Username {
			t.Errorf("Expected %v, got %v", i.Username, user.Username.String())
		}

		if !user.VerifyPassword(i.Password) {
			t.Error("The credentials of user is invalid")
		}
	})

	t.Run("should_return_error_if_user_not_exists_in_repository", func(t *testing.T) {
		i := auth.UserAuthenticatorInput{
			Username: "Jonathan James",
			Password: "1234",
		}

		_, err := u.Authenticate(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrAuthenticationFailed) {
			t.Errorf("The error must be %v, got %v", auth.ErrAuthenticationFailed, err)
		}
	})

	t.Run("should_return_error_if_password_is_invalid", func(t *testing.T) {
		i := auth.UserAuthenticatorInput{
			Username: "Jonathan",
			Password: "1234343",
		}

		_, err := u.Authenticate(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrAuthenticationFailed) {
			t.Errorf("The error must be %v, got %v", auth.ErrAuthenticationFailed, err)
		}
	})
}

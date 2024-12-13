package auth_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)

func TestUserAutenticator(t *testing.T) {
	user := auth.User{
		Username: "Jonathan",
		Password: "1234",
	}

	repo := auth.NewInmemRepository(user)

	t.Run("should_autenticate_user", func(t *testing.T) {
		u := auth.NewUserAutenticator(repo)
		i := auth.UserAutenticatorInput{
			Username: "Jonathan",
			Password: "1234",
		}

		_, err := u.Autenticate(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_autenticate_with_credendentials", func(t *testing.T) {
		u := auth.NewUserAutenticator(repo)
		i := auth.UserAutenticatorInput{Username: "Jonathan", Password: "1234"}

		o, err := u.Autenticate(i)

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

	t.Run("should_verify_if_credentials_are_corrects", func(t *testing.T) {
		u := auth.NewUserAutenticator(repo)
		i := auth.UserAutenticatorInput{
			Username: "Jonathan",
			Password: "1234",
		}

		o, err := u.Autenticate(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.Username != i.Username {
			t.Errorf("Expected %v, got %v", i.Username, o.Username)
		}

		if o.Password != i.Password {
			t.Errorf("Expected %v, got %v", i.Password, o.Password)
		}
	})

	t.Run("should_return_error_if_user_not_exists_in_repository", func(t *testing.T) {
		u := auth.NewUserAutenticator(repo)
		i := auth.UserAutenticatorInput{
			Username: "Jonathan James",
			Password: "1234",
		}

		_, err := u.Autenticate(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrAutenticateFailed) {
			t.Errorf("The error must be %v, got %v", auth.ErrAutenticateFailed, err)
		}
	})

	t.Run("should_return_error_if_password_is_invalid", func(t *testing.T) {
		u := auth.NewUserAutenticator(repo)
		i := auth.UserAutenticatorInput{
			Username: "Jonathan",
			Password: "1234343",
		}

		_, err := u.Autenticate(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrPasswordInvalid) {
			t.Errorf("The error must be %v, got %v", auth.ErrPasswordInvalid, err)
		}
	})

}

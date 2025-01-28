package auth_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)

func TestUserPasswordResetter(t *testing.T) {
	user := auth.User{Email: "kate@gmail.com", Password: auth.MustNewPassword("kate1234")}
	repo := auth.NewInmemRepository(user)
	u := auth.NewUserPasswordResetter(repo)
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
			t.Error("should update password")
		}
	})

	t.Run("should_send_new_password_to_email", func(t *testing.T) {
		
	})
}

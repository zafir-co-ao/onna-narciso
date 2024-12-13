package auth_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)


func TestUserCreator(t *testing.T) {
	repo := auth.NewInmemRepository()
	u := auth.NewUserCreator(repo)

	t.Run("should_create_a_user", func(t *testing.T) {

		_, err := u.Create()

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_save_user_in_repository", func(t *testing.T) {
		o, err := u.Create()

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		u, err := repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("Should return a user from repository got %v", err)
		}

		if u.ID.String() != o.ID {
			t.Errorf("The user id must be equal to %v, got %v", o.ID, u.ID.String())
		}
	})
}

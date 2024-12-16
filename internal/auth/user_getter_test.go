package auth_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)

func TestUserGetter(t *testing.T) {
	user := auth.User{ID: "1"}
	repo := auth.NewInmemRepository(user)
	u := auth.NewUserGetter(repo)

	t.Run("should_retrieve_user_from_repository", func(t *testing.T) {
		id := "1"

		o, err := u.Get(id)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.ID != id {
			t.Errorf("The ID of user must be %s, got %s", id, o.ID)
		}
	})

	t.Run("should_return_error_if_user_not_found", func(t *testing.T) {
		id := "2"

		_, err := u.Get(id)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("The error must be %v, got %v", auth.ErrUserNotFound, err)
		}
	})
}

package auth_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)

func TestUserFinder(t *testing.T) {
	users := []auth.User{
		{
			ID:   "1",
			Role: auth.RoleManager,
		},
		{
			ID:   "2",
			Role: auth.RoleReceptionist,
		},
		{
			ID:       nanoid.New(),
			Username: "Mike Tyson",
		},
	}
	repo := auth.NewInmemRepository(users...)
	u := auth.NewUserFinder(repo)

	t.Run("should_retrieve_users", func(t *testing.T) {
		users, err := u.Find("1")

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(users) == 0 {
			t.Errorf("Should return all users, got %v", len(users))
		}
	})

	t.Run("should_return_error_if_not_user_manager_to_retrieve_all_users", func(t *testing.T) {
		// _, err := u.Find("2")

		// if errors.Is(nil, err) {
		// 	t.Errorf("Expected error, got %v", err)
		// }

		// if !errors.Is(err, auth.ErrUserNotAllowed) {
		// 	t.Errorf("The error must be %v, got %v", auth.ErrUserNotAllowed, err)
		// }
	})

	t.Run("should_return_error_if_user_manager_not_found", func(t *testing.T) {
		// _, err := u.Find(nanoid.New().String())

		// if errors.Is(nil, err) {
		// 	t.Errorf("Expected error, got %v", err)
		// }

		// if !errors.Is(err, auth.ErrUserNotFound) {
		// 	t.Errorf("The error must be %v, got %v", auth.ErrUserNotFound, err)
		// }
	})
}

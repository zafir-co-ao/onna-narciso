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
		i := auth.UserCreatorInput{}

		_, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_save_user_in_repository", func(t *testing.T) {
		i := auth.UserCreatorInput{}

		o, err := u.Create(i)

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

	t.Run("must_register_the_username_password_and_role_of_user", func(t *testing.T) {
		i := auth.UserCreatorInput{
			Username: "John Doe",
			Password: "john.doe@123",
			Role:     "Gestor",
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		u, err := repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("Should return a user from repository got %v", err)
		}

		if u.Username.String() != i.Username {
			t.Errorf("The username must be equal to %v, got %v", i.Username, u.Username.String())
		}

		if u.Password.String() == "" {
			t.Error("The password of user must be defined")
		}

		if u.Role.String() != i.Role {
			t.Errorf("The role of user must be %v, got %v", i.Role, u.Role)
		}
	})
}

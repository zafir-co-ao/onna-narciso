package auth_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)

func TestUserCreator(t *testing.T) {
	users := []auth.User{
		{ID: "1", Username: "Paola Oliveira", Role: auth.RoleManager},
		{ID: "2", Username: "Elon Musk", Role: auth.RoleReceptionist},
	}

	bus := event.NewEventBus()
	repo := auth.NewInmemRepository(users...)
	u := auth.NewUserCreator(repo, bus)

	t.Run("should_create_a_user", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:   "1",
			Username: "Mike Tyson",
			Password: "somepassword",
			Role:     auth.RoleCustomer.String(),
		}

		_, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_save_user_in_repository", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:   "1",
			Username: "Erling Haaland",
			Password: "erlingpassword",
			Role:     auth.RoleManager.String(),
		}

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
			UserID:   "1",
			Username: "John Doe",
			Password: "john.doe@123",
			Role:     auth.RoleManager.String(),
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

	t.Run("should_protect_the_password_of_user", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:   "1",
			Username: "Joana Doe",
			Password: "joana.doe@123",
			Role:     auth.RoleReceptionist.String(),
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		u, err := repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("Should return a user from repository got %v", err)
		}

		if u.Password.String() == i.Password {
			t.Errorf("Should register the hash of password, got %v", u.Password.String())
		}

		if !u.VerifyPassword(i.Password) {
			t.Error("Should verify the password of user")
		}
	})

	t.Run("should_publish_event_when_user_is_created", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:   "1",
			Username: "Martin Fowler",
			Password: "martin.fowler@123",
			Role:     auth.RoleCustomer.String(),
		}

		isPublished := false

		var h event.HandlerFunc = func(e event.Event) {
			if e.Name() == auth.EventUserCreated {
				isPublished = true
			}
		}

		bus.Subscribe(auth.EventUserCreated, h)

		_, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("The %v must be published", auth.EventUserCreated)
		}
	})

	t.Run("should_return_error_if_role_of_user_is_not_allowed", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:   "1",
			Username: "Robert C. Martin",
			Password: "robert.martin@0000",
			Role:     "Role",
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected an error got, %v", err)
		}

		if !errors.Is(err, auth.ErrRoleNotAllowed) {
			t.Errorf("The error must be %v, got %v", auth.ErrRoleNotAllowed, err)
		}
	})

	t.Run("should_return_error_if_username_not_provided", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:   "1",
			Username: "",
			Password: "cliente@9999",
			Role:     auth.RoleCustomer.String(),
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected an error got, %v", err)
		}

		if !errors.Is(err, auth.ErrEmptyUsername) {
			t.Errorf("The error must be %v, got %v", auth.ErrEmptyUsername, err)
		}
	})

	t.Run("should_return_error_if_password_not_provided", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:   "1",
			Username: "Gustavo Lima",
			Password: "",
			Role:     auth.RoleManager.String(),
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected an error got, %v", err)
		}

		if !errors.Is(err, auth.ErrEmptyPassword) {
			t.Errorf("The error must be %v, got %v", auth.ErrEmptyPassword, err)
		}
	})

	t.Run("should_return_error_if_is_not_manager_to_create_a_new_user", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:   "2",
			Username: "Rafa Nadal",
			Password: "rafanadaltenis",
			Role:     auth.RoleManager.String(),
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected an error got, %v", err)
		}

		if !errors.Is(err, auth.ErrUserNotAllowed) {
			t.Errorf("The error must be %v, got %v", auth.ErrUserNotAllowed, err)
		}
	})

	t.Run("should_return_error_if_user_not_found", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:   "3",
			Username: "Paola Oliveira",
			Password: "paolaoliveira",
			Role:     auth.RoleManager.String(),
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected an error got, %v", err)
		}

		if !errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("The error must be %v, got %v", auth.ErrUserNotFound, err)
		}
	})

	t.Run("should_return_error_if_username_not_is_unique", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:   "1",
			Username: "Paola Oliveira",
			Password: "paolaoliveira",
			Role:     auth.RoleManager.String(),
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected an error got, %v", err)
		}

		if !errors.Is(err, auth.ErrOnlyUniqueUsername) {
			t.Errorf("The error must be %v, got %v", auth.ErrOnlyUniqueUsername, err)
		}
	})
}

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
		{
			ID:          "3",
			Username:    "Robertson Konnely",
			Email:       "bod@gmail.com",
			PhoneNumber: "923459876",
			Role:        auth.RoleReceptionist,
		},
	}

	bus := event.NewEventBus()
	repo := auth.NewInmemRepository(users...)
	u := auth.NewUserCreator(repo, bus)

	t.Run("should_create_an_user", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:      "1",
			Username:    "Mike Tyson",
			Email:       "miketyson@gmail.com",
			PhoneNumber: "982342312",
			Password:    "somepassword",
			Role:        auth.RoleCustomer.String(),
		}

		_, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_save_user_in_repository", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:      "1",
			Username:    "Erling Haaland",
			Email:       "erling@gmail.com",
			PhoneNumber: "934231234",
			Password:    "erlingpassword",
			Role:        auth.RoleManager.String(),
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		user, err := repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("Should return a user from repository got %v", err)
		}

		if user.ID.String() != o.ID {
			t.Errorf("The user id must be equal to %v, got %v", o.ID, user.ID.String())
		}
	})

	t.Run("must_register_the_username_password_and_role_of_user", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:      "1",
			Username:    "John Doe",
			Email:       "john@gmail.com",
			PhoneNumber: "934123456",
			Password:    "john.doe@123",
			Role:        auth.RoleManager.String(),
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		user, err := repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("Should return a user from repository got %v", err)
		}

		if user.Username.String() != i.Username {
			t.Errorf("The username must be equal to %v, got %v", i.Username, user.Username.String())
		}

		if user.Password.String() == "" {
			t.Error("The password of user must be defined")
		}

		if user.Role.String() != i.Role {
			t.Errorf("The role of user must be %v, got %v", i.Role, user.Role)
		}
	})

	t.Run("must_register_email_and_phonenumber", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:      "1",
			Username:    "Joana Becket",
			Password:    "john.doe@123",
			Email:       "joana@gmail.com",
			PhoneNumber: "932345412",
			Role:        auth.RoleManager.String(),
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		user, err := repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("Should return a user from repository got %v", err)
		}

		if user.Email.String() != i.Email {
			t.Errorf("Email must be equal to %v, got %v", i.Email, user.Email)
		}

		if user.PhoneNumber.String() != i.PhoneNumber {
			t.Errorf("Email must be equal to %v, got %v", i.PhoneNumber, user.PhoneNumber)
		}
	})

	t.Run("should_protect_the_password_of_user", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:      "1",
			Username:    "Joana Doe",
			Email:       "joana2@gmail.com",
			PhoneNumber: "98765423",
			Password:    "joana.doe@123",
			Role:        auth.RoleReceptionist.String(),
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		user, err := repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("Should return a user from repository got %v", err)
		}

		if user.Password.String() == i.Password {
			t.Errorf("Should register the hash of password, got %v", user.Password.String())
		}

		if !user.VerifyPassword(i.Password) {
			t.Error("Should verify the password of user")
		}
	})

	t.Run("should_publish_event_when_user_is_created", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:      "1",
			Username:    "Martin Fowler",
			Email:       "martin@gmail.com",
			PhoneNumber: "931235412",
			Password:    "martin.fowler@123",
			Role:        auth.RoleCustomer.String(),
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
			UserID:      "1",
			Username:    "Robert C. Martin",
			Email:       "bod3@gmail.com",
			PhoneNumber: "923459877",
			Password:    "robert.martin@0000",
			Role:        "Role",
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

	t.Run("showld_return_error_if_email_not_provided", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:      "1",
			Username:    "James Trubin",
			Password:    "john.doe@123",
			Email:       "",
			PhoneNumber: "932345412",
			Role:        auth.RoleManager.String(),
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected an error got, %v", err)
		}

		if !errors.Is(err, auth.ErrEmptyEmail) {
			t.Errorf("The error must be %v, got %v", auth.ErrEmptyEmail, err)
		}
	})

	t.Run("should_return_error_if_phonenumber_not_provided", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:      "1",
			Username:    "James Trubin",
			Password:    "john.doe@123",
			Email:       "james@gmail.com",
			PhoneNumber: "",
			Role:        auth.RoleManager.String(),
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected an error got, %v", err)
		}

		if !errors.Is(err, auth.ErrEmptyPhoneNumber) {
			t.Errorf("The error must be %v, got %v", auth.ErrEmptyPhoneNumber, err)
		}
	})

	t.Run("should_return_error_if_email_is_in_invalid_format", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:      "1",
			Username:    "James Trubin",
			Password:    "john.doe@123",
			Email:       "james@gmail",
			PhoneNumber: "9124313432",
			Role:        auth.RoleManager.String(),
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected an error got, %v", err)
		}

		if !errors.Is(err, auth.ErrInvalidEmailFormat) {
			t.Errorf("The error must be %v, got %v", auth.ErrInvalidEmailFormat, err)
		}
	})

	t.Run("should_return_error_if_password_not_provided", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:      "1",
			Username:    "Gustavo Lima",
			Email:       "gustavo@gmail.com",
			PhoneNumber: "934123498",
			Password:    "",
			Role:        auth.RoleManager.String(),
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
			UserID:   "3U",
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

	t.Run("should_return_error_if_user_email_is_not_unique", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:      "1",
			Username:    "Robert C. Martin",
			Email:       "bod@gmail.com",
			PhoneNumber: "923459876",
			Password:    "robert.martin@0000",
			Role:        auth.RoleManager.String(),
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected an error got, %v", err)
		}

		if !errors.Is(err, auth.ErrOnlyUniqueEmail) {
			t.Errorf("The error must be %v, got %v", auth.ErrOnlyUniqueEmail, err)
		}
	})

	t.Run("should_return_error_if_phonenumber_is_not_unique", func(t *testing.T) {
		i := auth.UserCreatorInput{
			UserID:      "1",
			Username:    "Yolanda Simões",
			Email:       "yola@gmail.com",
			PhoneNumber: "923459876",
			Password:    "yola@0000",
			Role:        auth.RoleManager.String(),
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected an error got, %v", err)
		}

		if !errors.Is(err, auth.ErrOnlyUniquePhoneNumber) {
			t.Errorf("The error must be %v, got %v", auth.ErrOnlyUniquePhoneNumber, err)
		}
	})
}

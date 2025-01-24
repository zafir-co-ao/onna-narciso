package auth_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)

func TestUserUpdater(t *testing.T) {
	users := []auth.User{
		{
			ID:          "1",
			Username:    "james",
			Email:       "james@gmail.com",
			PhoneNumber: "9431234567",
			Role:        auth.RoleManager,
		},
		{
			ID:          "2",
			Username:    "James Kluivert",
			Email:       "kluivert@gmail.com",
			PhoneNumber: "923123421",
			Password:    "1234",
			Role:        auth.RoleReceptionist,
		},
		{
			ID:   "3",
			Role: auth.RoleCustomer,
		},
	}
	repo := auth.NewInmemRepository(users...)
	bus := event.NewEventBus()
	u := auth.NewUserUpdater(repo, bus)

	t.Run("should_update_user", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "Ana",
			Email:       "ana@gmail.com",
			PhoneNumber: "932123432",
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("should_retrieve_user_in_repository", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "Arthur",
			Email:       "arthur@gmail.com",
			PhoneNumber: "983124312",
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("Should find a user in repository, got %v", err)
		}
	})

	t.Run("should_update_username", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "GomesDoyle",
			Email:       "gomes@gmail.com",
			PhoneNumber: "912345678",
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		user, err := repo.FindByID(nanoid.ID(i.UserID))
		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("Should find a user in repository, got %v", err)
		}

		if user.Username.String() != i.Username {
			t.Errorf("The expected is %v, got %v", i.Username, user.Username.String())
		}
	})

	t.Run("should_update_email", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "doyleDan",
			Email:       "gomes3@gmail.com",
			PhoneNumber: "945321234",
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		user, err := repo.FindByID(nanoid.ID(i.UserID))
		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("Should find a user in repository, got %v", err)
		}

		if user.Email.String() != i.Email {
			t.Errorf("The expected is %v, got %v", i.Email, user.Email.String())
		}
	})

	t.Run("should_update_phone_number", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "sandra",
			Email:       "234s@gmail.com",
			PhoneNumber: "934123456",
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		user, err := repo.FindByID(nanoid.ID(i.UserID))
		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("Should find a user in repository, got %v", err)
		}

		if user.PhoneNumber.String() != i.PhoneNumber {
			t.Errorf("The expected is %v, got %v", i.PhoneNumber, user.PhoneNumber.String())
		}
	})

	t.Run("should_update_role", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "Darling",
			Email:       "darling@gmail.com",
			PhoneNumber: "934123457",
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		user, err := repo.FindByID(nanoid.ID(i.UserID))
		if errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("Should find a user in repository, got %v", err)
		}

		if user.Role.String() != i.Role {
			t.Errorf("The expected is %v, got %v", i.Role, user.Role.String())
		}
	})

	t.Run("should_publish_domain_event_when_user_updated", func(t *testing.T) {
		isPublished := false

		var h event.HandlerFunc = func(e event.Event) {
			if e.Name() == auth.EventUserUpdated {
				isPublished = true
			}
		}

		bus.Subscribe(auth.EventUserUpdated, h)

		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "Darling Parker",
			Email:       "parker@gmail.com",
			PhoneNumber: "934123417",
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("The %s must be published", auth.EventUserUpdated)
		}
	})

	t.Run("should_return_error_if_user_exists_in_repsository", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "1D",
			Username:    "António",
			Email:       "gomes@gmail.com",
			PhoneNumber: "934123456",
			
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("The error must be %v, got %v", auth.ErrUserNotFound, err)
		}
	})

	t.Run("should_return_error_if_username_not_provided", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "",
			Email:       "gomes@gmail.com",
			PhoneNumber: "934123456",
			
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrEmptyUsername) {
			t.Errorf("The error must be %v, got %v", auth.ErrEmptyUsername, err)
		}
	})

	t.Run("should_return_error_if_email_not_provided", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "Ana",
			Email:       "",
			PhoneNumber: "934123456",
			
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrEmptyEmail) {
			t.Errorf("The error must be %v, got %v", auth.ErrEmptyEmail, err)
		}
	})

	t.Run("should_return_error_if_phonenumber_not_provided", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "Ana",
			Email:       "ana@gmail.com",
			PhoneNumber: "",
			
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrEmptyPhoneNumber) {
			t.Errorf("The error must be %v, got %v", auth.ErrEmptyPhoneNumber, err)
		}
	})

	t.Run("should_return_error_if_role_of_user_is_not_allowed", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "Ana",
			Email:       "ana@gmail.com",
			PhoneNumber: "9431234547",
			
			Role:        "Some",
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrRoleNotAllowed) {
			t.Errorf("The error must be %v, got %v", auth.ErrRoleNotAllowed, err)
		}
	})

	t.Run("should_return_error_if_is_not_role_manager_to_update_user", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "3",
			UserID:      "2",
			Username:    "Ana",
			Email:       "ana@gmail.com",
			PhoneNumber: "9431234567",
			
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrUserNotAllowed) {
			t.Errorf("The error must be %v, got %v", auth.ErrUserNotAllowed, err)
		}
	})

	t.Run("should_return_error_if_manager_not_exists_in_repository", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "3D",
			UserID:      "2",
			Username:    "Ana",
			Email:       "ana@gmail.com",
			PhoneNumber: "9431234567",
			
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrUserNotFound) {
			t.Errorf("The error must be %v, got %v", auth.ErrUserNotFound, err)
		}
	})

	t.Run("should_return_error_if_username_is_not_unique", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "james",
			Email:       "james234@gmail.com",
			PhoneNumber: "9431234567",
			
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrOnlyUniqueUsername) {
			t.Errorf("The error must be %v, got %v", auth.ErrOnlyUniqueUsername, err)
		}
	})

	t.Run("should_return_error_if_email_is_not_unique", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "Ana",
			Email:       "james@gmail.com",
			PhoneNumber: "9431234567",
			
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrOnlyUniqueEmail) {
			t.Errorf("The error must be %v, got %v", auth.ErrOnlyUniqueEmail, err)
		}
	})

	t.Run("should_return_error_if_phonenumber_is_not_unique", func(t *testing.T) {
		i := auth.UserUpdaterInput{
			ManagerID:   "1",
			UserID:      "2",
			Username:    "Ana",
			Email:       "james2@gmail.com",
			PhoneNumber: "9431234567",
			
			Role:        auth.RoleReceptionist.String(),
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("expected error, got %v", err)
		}

		if !errors.Is(err, auth.ErrOnlyUniquePhoneNumber) {
			t.Errorf("The error must be %v, got %v", auth.ErrOnlyUniquePhoneNumber, err)
		}
	})
}

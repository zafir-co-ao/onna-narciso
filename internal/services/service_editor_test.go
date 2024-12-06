package services_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/services/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/services/price"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

func TestServiceEdit(t *testing.T) {
	bus := event.NewEventBus()
	repo := inmem.NewServiceRepository()

	s := services.Service{
		ID:       nanoid.ID("1"),
		Name:     "Manicure",
		Price:    "1000",
		Duration: 60,
	}

	_ = repo.Save(s)

	u := services.NewServiceEditor(repo, bus)

	t.Run("should_recovery_service_with_id", func(t *testing.T) {
		i := services.ServiceEditorInput{
			ID:       s.ID.String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 90,
		}

		err := u.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(s.ID)
		if errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("should be find service in repository")
		}

		if s.ID.String() != i.ID {
			t.Errorf("expected %v, got %v", i.ID, s.ID.String())
		}

	})

	t.Run("should_edit_name_of_service", func(t *testing.T) {
		i := services.ServiceEditorInput{
			ID:       s.ID.String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 90,
		}

		err := u.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(s.ID)
		if errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("should be find service in repository")
		}

		if s.Name.String() != i.Name {
			t.Errorf("expected %v, got %v", i.Name, s.Name.String())
		}
	})

	t.Run("should_edit_price_of_service", func(t *testing.T) {
		i := services.ServiceEditorInput{
			ID:       s.ID.String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 90,
		}

		err := u.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(s.ID)
		if errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("should be find service in repository")
		}

		if string(s.Price) != i.Price {
			t.Errorf("expected %v, got %v", i.Price, string(s.Price))
		}
	})

	t.Run("should_edit_description_of_service", func(t *testing.T) {
		i := services.ServiceEditorInput{
			ID:          s.ID.String(),
			Name:        "Manicure e Pedicure",
			Price:       "1500",
			Description: "Com Gelinho no Pé",
			Duration:    90,
		}

		err := u.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(s.ID)
		if errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("should be find service in repository")
		}

		if string(s.Description) != i.Description {
			t.Errorf("expected %v, got %v", i.Description, string(s.Description))
		}
	})

	t.Run("should_edit_duration_of_service", func(t *testing.T) {
		i := services.ServiceEditorInput{
			ID:       s.ID.String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 90,
		}

		err := u.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(s.ID)
		if errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("should be find service in repository")
		}

		if s.Duration.Value() != i.Duration {
			t.Errorf("expected %v, got %v", i.Duration, s.Duration.Value())
		}
	})

	t.Run("should_publish_the_domain_event_when_service_is_edited", func(t *testing.T) {
		isPublished := false
		var h event.HandlerFunc = func(e event.Event) {
			if e.Name() == services.EventServiceEdited {
				isPublished = true
			}
		}

		i := services.ServiceEditorInput{
			ID:       s.ID.String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 90,
		}

		bus.Subscribe(services.EventServiceEdited, h)

		err := u.Edit(i)
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("The %v should be published", services.EventServiceEdited)
		}
	})

	t.Run("should_return_error_if_name_is_empty", func(t *testing.T) {
		i := services.ServiceEditorInput{
			ID:       s.ID.String(),
			Price:    "1500",
			Duration: 90,
		}

		err := u.Edit(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, name.ErrEmptyName) {
			t.Errorf("The error must be %v, got %v", name.ErrEmptyName, err)
		}
	})

	t.Run("should_return_error_if_price_is_empty", func(t *testing.T) {
		i := services.ServiceEditorInput{
			ID:       s.ID.String(),
			Name:     "Manicure e Pedicure",
			Duration: 90,
		}

		err := u.Edit(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, price.ErrInvalidPrice) {
			t.Errorf("The error must be %v, got %v", price.ErrInvalidPrice, err)
		}
	})

	t.Run("should_return_error_if_duration_is_invalid", func(t *testing.T) {
		i := services.ServiceEditorInput{
			ID:       s.ID.String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: -120,
		}

		err := u.Edit(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, duration.ErrInvalidDuration) {
			t.Errorf("The error must be %v, got %v", duration.ErrInvalidDuration, err)
		}
	})

	t.Run("should_return_error_if_service_not_exit_in_repository", func(t *testing.T) {
		i := services.ServiceEditorInput{
			ID:       nanoid.New().String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 120,
		}

		err := u.Edit(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("The error must be %v, got %v", services.ErrServiceNotFound, err)
		}
	})
}

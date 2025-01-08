package services_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

func TestServiceUpdate(t *testing.T) {

	s := []services.Service{
		{
			ID:       nanoid.ID("1"),
			Name:     "Manicure",
			Price:    "1000",
			Duration: 60,
		},
		{
			ID:       nanoid.ID("2"),
			Name:     "Manicure",
			Price:    "1000",
			Duration: 60,
		},
	}

	bus := event.NewEventBus()
	repo := services.NewInmemRepository(s...)
	u := services.NewServiceUpdater(repo, bus)

	t.Run("should_retrieve_service_with_id", func(t *testing.T) {
		i := services.ServiceUpdaterInput{
			ID:       s[0].ID.String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 90,
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(s[0].ID)
		if errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("should find service in repository")
		}

		if s.ID.String() != i.ID {
			t.Errorf("expected %v, got %v", i.ID, s.ID.String())
		}
	})

	t.Run("should_update_name_of_service", func(t *testing.T) {
		i := services.ServiceUpdaterInput{
			ID:       s[0].ID.String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 90,
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(s[0].ID)
		if errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("should find service in repository")
		}

		if s.Name.String() != i.Name {
			t.Errorf("expected %v, got %v", i.Name, s.Name.String())
		}
	})

	t.Run("should_update_price_of_service", func(t *testing.T) {
		i := services.ServiceUpdaterInput{
			ID:       s[0].ID.String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 90,
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(s[0].ID)
		if errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("should find service in repository")
		}

		if string(s.Price) != i.Price {
			t.Errorf("expected %v, got %v", i.Price, string(s.Price))
		}
	})

	t.Run("should_update_description_of_service", func(t *testing.T) {
		i := services.ServiceUpdaterInput{
			ID:          s[0].ID.String(),
			Name:        "Manicure e Pedicure",
			Price:       "1500",
			Description: "Com Gelinho no Pé",
			Duration:    90,
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(s[0].ID)
		if errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("should be find service in repository")
		}

		if string(s.Description) != i.Description {
			t.Errorf("expected %v, got %v", i.Description, string(s.Description))
		}
	})

	t.Run("should_update_duration_of_service", func(t *testing.T) {
		i := services.ServiceUpdaterInput{
			ID:       s[0].ID.String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 90,
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(s[0].ID)
		if errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("should find service in repository, got %v", services.ErrServiceNotFound)
		}

		if s.Duration.Value() != i.Duration {
			t.Errorf("expected %v, got %v", i.Duration, s.Duration.Value())
		}
	})

	t.Run("should_update_discount_of_service", func(t *testing.T) {
		i := services.ServiceUpdaterInput{
			ID:       s[0].ID.String(),
			Name:     "Massagem",
			Price:    "16000",
			Duration: 120,
			Discount: "50",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(s[0].ID)
		if errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("should find service in repository, got %v", services.ErrServiceNotFound)
		}

		if string(s.Discount) != i.Discount {
			t.Errorf("The discount of service must be equal to %v, got %v", i.Discount, string(s.Discount))
		}
	})

	t.Run("should_publish_the_event_price_updated_when_price_is_updated", func(t *testing.T) {
		isPublished := false
		var h event.HandlerFunc = func(e event.Event) {
			if e.Name() == services.EventPriceUpdated {
				isPublished = true
			}
		}

		bus.Subscribe(services.EventPriceUpdated, h)

		i := services.ServiceUpdaterInput{
			ID:       s[1].ID.String(),
			Name:     "Manicure e Pedicure",
			Price:    "1200",
			Duration: 60,
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("The %v should be published, %v", services.EventPriceUpdated, isPublished)
		}
	})

	t.Run("should_return_error_if_name_is_empty", func(t *testing.T) {
		i := services.ServiceUpdaterInput{
			ID:       s[0].ID.String(),
			Price:    "1500",
			Duration: 90,
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, name.ErrEmptyName) {
			t.Errorf("The error must be %v, got %v", name.ErrEmptyName, err)
		}
	})

	t.Run("should_return_error_if_price_is_empty", func(t *testing.T) {
		i := services.ServiceUpdaterInput{
			ID:       s[0].ID.String(),
			Name:     "Manicure e Pedicure",
			Duration: 90,
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, services.ErrInvalidPrice) {
			t.Errorf("The error must be %v, got %v", services.ErrInvalidPrice, err)
		}
	})

	t.Run("should_return_error_if_duration_is_invalid", func(t *testing.T) {
		i := services.ServiceUpdaterInput{
			ID:       s[0].ID.String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: -120,
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, duration.ErrInvalidDuration) {
			t.Errorf("The error must be %v, got %v", duration.ErrInvalidDuration, err)
		}
	})

	t.Run("should_return_error_if_service_not_exists_in_repository", func(t *testing.T) {
		i := services.ServiceUpdaterInput{
			ID:       nanoid.New().String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 120,
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("The error must be %v, got %v", services.ErrServiceNotFound, err)
		}
	})

	t.Run("should_return_error_if_discount_is_invalid", func(t *testing.T) {
		i := services.ServiceUpdaterInput{
			ID:       s[0].ID.String(),
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 60,
			Discount: "10000",
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, services.ErrDiscountNotAllowed) {
			t.Errorf("The error must be %v, got %v", services.ErrDiscountNotAllowed, err)
		}
	})
}

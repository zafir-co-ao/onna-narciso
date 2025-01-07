package hr_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/hr"
	"github.com/zafir-co-ao/onna-narciso/internal/hr/stubs"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

func TestProfesssionalUpdater(t *testing.T) {
	professionals := []hr.Professional{
		{
			ID:   nanoid.ID("1"),
			Name: "Jing Woo",
		},
	}
	repo := hr.NewInmemProfessionalRepository(professionals...)
	sacl := stubs.NewServicesServiceACL()
	bus := event.NewEventBus()
	u := hr.NewProfessionalUpdater(repo, sacl, bus)

	t.Run("should_retrieve_professional", func(t *testing.T) {
		i := hr.ProfessionalUpdaterInput{
			ID:          "1",
			Name:        "Jennifer",
			ServicesIDs: []string{"1"},
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		o, err := repo.FindByID(nanoid.ID(i.ID))
		if errors.Is(err, hr.ErrProfessionalNotFound) {
			t.Errorf("Should find a professional in repository, got %v", err)
		}

		if o.ID.String() != i.ID {
			t.Errorf("expected %v, got %v", i.ID, o.ID.String())
		}
	})

	t.Run("should_update_professional_name", func(t *testing.T) {
		i := hr.ProfessionalUpdaterInput{
			ID:          "1",
			Name:        "Jonathan Paulino",
			ServicesIDs: []string{"1"},
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		o, err := repo.FindByID(nanoid.ID(i.ID))
		if errors.Is(err, hr.ErrProfessionalNotFound) {
			t.Errorf("Should find a professional in repository, got %v", err)
		}

		if o.Name.String() != i.Name {
			t.Errorf("expected %v, got %v", i.Name, o.Name.String())
		}
	})

	t.Run("should_update_professional_services", func(t *testing.T) {
		i := hr.ProfessionalUpdaterInput{
			ID:          "1",
			Name:        "Jonathan Paulino",
			ServicesIDs: []string{"1", "2", "3"},
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		o, err := repo.FindByID(nanoid.ID(i.ID))
		if errors.Is(err, hr.ErrProfessionalNotFound) {
			t.Errorf("Should find a professional in repository, got %v", err)
		}

		if len(o.Services) != 3 {
			t.Errorf("expected %v, got %v", 3, len(o.Services))
		}
	})

	t.Run("should_return_error_when_professional_not_exists", func(t *testing.T) {
		i := hr.ProfessionalUpdaterInput{
			ID:          "1P",
			Name:        "Jing Woo",
			ServicesIDs: []string{"1", "2", "3"},
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, hr.ErrProfessionalNotFound) {
			t.Errorf("The error must be %v, got %v", hr.ErrProfessionalNotFound, err)
		}
	})

	t.Run("should_return_error_when_name_empty", func(t *testing.T) {
		i := hr.ProfessionalUpdaterInput{
			ID:          "1",
			Name:        "",
			ServicesIDs: []string{"1", "2", "3"},
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, name.ErrEmptyName) {
			t.Errorf("The error must be %v, got %v", name.ErrEmptyName, err)
		}
	})

	t.Run("should_return_error_when_service_not_exists", func(t *testing.T) {
		i := hr.ProfessionalUpdaterInput{
			ID:          "1",
			Name:        "Sky Aline",
			ServicesIDs: []string{"1", "2", "3SD"},
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, hr.ErrServiceNotFound) {
			t.Errorf("The error must be %v, got %v", hr.ErrServiceNotFound, err)
		}
	})

	t.Run("should_publish_domain_event_professional_updated", func(t *testing.T) {
		i := hr.ProfessionalUpdaterInput{
			ID:          "1",
			Name:        "Jonathan Paulino",
			ServicesIDs: []string{"1"},
		}

		isPublished := false

		var h = event.HandlerFunc(func(e event.Event) {
			if e.Name() == hr.EventProfessionalUpdated {
				isPublished = true
			}
		})

		bus.Subscribe(hr.EventProfessionalUpdated, h)

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("Should publish the event %v", hr.EventProfessionalUpdated)
		}
	})
}

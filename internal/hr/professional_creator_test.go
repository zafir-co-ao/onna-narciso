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

func TestProfessionalCreator(t *testing.T) {
	bus := event.NewEventBus()
	sacl := stubs.NewServicesServiceACL()
	repo := hr.NewProfessionalRepository()
	u := hr.NewProfessionalCreator(repo, sacl, bus)

	t.Run("should_create_a_professional", func(t *testing.T) {
		i := hr.ProfessionalCreatorInput{Name: "Jonathan"}

		_, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_store_professional_in_repository", func(t *testing.T) {
		i := hr.ProfessionalCreatorInput{Name: "Jonathan"}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		_, err = repo.FindByID(nanoid.ID(o.ID))

		if errors.Is(err, hr.ErrProfessionalNotFound) {
			t.Errorf("Should return a professional from repository")
		}
	})

	t.Run("should_create_professional_with_the_name", func(t *testing.T) {
		i := hr.ProfessionalCreatorInput{Name: "Jonathan"}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		p, err := repo.FindByID(nanoid.ID(o.ID))

		if errors.Is(err, hr.ErrProfessionalNotFound) {
			t.Errorf("Should return a professional from repository, got %v", err)
		}

		if p.Name.String() != i.Name {
			t.Errorf("The expected name is %v, got %v", i.Name, p.Name.String())
		}
	})

	t.Run("should_create_professional_with_respectives_services", func(t *testing.T) {
		i := hr.ProfessionalCreatorInput{
			Name:        "Jonathan",
			ServicesIDs: []string{"1", "2"},
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		p, err := repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(err, hr.ErrProfessionalNotFound) {
			t.Errorf("Should return a professional, got %v", hr.ErrProfessionalNotFound)
		}

		if len(p.Services) != 2 {
			t.Errorf("Should has %v in repository, got %v", 2, p.Services)
		}

		if !p.HasService(nanoid.ID(i.ServicesIDs[0])) {
			t.Errorf("The service expected is %v, got %v", i.ServicesIDs[0], p.Services[0].ID.String())
		}

		if !p.HasService(nanoid.ID(i.ServicesIDs[1])) {
			t.Errorf("The service expected is %v, got %v", i.ServicesIDs[1], p.Services[1].ID.String())
		}
	})

	t.Run("should_publish_the_domain_event_professional_created", func(t *testing.T) {
		isPublished := false

		var h event.HandlerFunc = func(e event.Event) {
			if e.Name() == hr.EventProfessionalCreated {
				isPublished = true
			}
		}

		bus.Subscribe(hr.EventProfessionalCreated, h)

		i := hr.ProfessionalCreatorInput{
			Name:        "Jonathan",
			ServicesIDs: []string{"1", "2"},
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		_, err = repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(err, hr.ErrProfessionalNotFound) {
			t.Errorf("Should return a professional, got %v", hr.ErrProfessionalNotFound)
		}

		if !isPublished {
			t.Errorf("Should publish the event %v", hr.EventProfessionalCreated)
		}
	})

	t.Run("should_return_error_if_name_is_empty", func(t *testing.T) {
		i := hr.ProfessionalCreatorInput{
			Name: "",
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, name.ErrEmptyName) {
			t.Errorf("The error must be %v, got %v", name.ErrEmptyName, err)
		}
	})

	t.Run("should_return_error_when_not_service", func(t *testing.T) {
		i := hr.ProfessionalCreatorInput{
			Name:        "Victor Kickoff",
			ServicesIDs: []string{"1", "2", "3"},
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, hr.ErrServiceNotFound) {
			t.Errorf("The error must be %v, got %v", hr.ErrServiceNotFound, err)
		}
	})
}

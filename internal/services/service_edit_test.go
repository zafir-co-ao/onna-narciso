package services_test

import (
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/services/adapters/inmem"
)

func TestServiceEdit(t *testing.T) {

	t.Run("should_recovery_service_with_id", func(t *testing.T) {

		bus := event.NewEventBus()
		repo := inmem.NewServiceRepository()
		u := services.NewServiceCreator(repo, bus)

		i := services.ServiceCreatorInput{
			Name:     "Manicure",
			Price:    "1000",
			Duration: 60,
		}

		o, err := u.Create(i)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		_, err = repo.FindByID(nanoid.ID(o.ID))

		if err != nil {
			t.Errorf("Should return a service from repository, got %v", err)

		}
	})

}

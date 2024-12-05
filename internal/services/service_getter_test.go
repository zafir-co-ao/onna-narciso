package services_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/services/adapters/inmem"
)

func TestServiceGetter(t *testing.T) {
	i := services.Service{
		ID:       "1",
		Name:     "Pedicure",
		Price:    "10000",
		Duration: 90,
	}

	repo := inmem.NewServiceRepository()
	u := services.NewServiceGetter(repo)

	_ = repo.Save(i)

	t.Run("should_retrieve_an_service_in_repository", func(t *testing.T) {
		id := "1"

		o, err := u.Get(id)
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.ID != id {
			t.Errorf("The service id %v must be equal to %v", id, o.ID)
		}
	})

	t.Run("should_return_error_when_service_not_found_in_repository", func(t *testing.T) {
		_, err := u.Get(nanoid.New().String())

		if errors.Is(nil, err) {
			t.Errorf("The getter service should return error, got %v", err)
		}

		if !errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("The error must be %v, got %v", services.ErrServiceNotFound, err)
		}
	})
}

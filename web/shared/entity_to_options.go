package shared

import (
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/shared/components"
)

func entityToOptions[T any](data []T, f func(t T) components.InputOption) []components.InputOption {
	var options []components.InputOption
	for _, d := range data {
		options = append(options, f(d))
	}
	return options
}

func ProfessionalsToOptions(professionals []scheduling.Professional) []components.InputOption {
	return entityToOptions(professionals, func(p scheduling.Professional) components.InputOption {
		return components.InputOption{p.Name.String(), p.ID.String()}
	})
}

func ServicesToOptions(services []scheduling.Service) []components.InputOption {
	return entityToOptions(services, func(s scheduling.Service) components.InputOption {
		return components.InputOption{s.Name.String(), s.ID.String()}
	})
}

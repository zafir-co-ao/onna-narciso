package components

import (
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

func entityToOptions[T any](data []T, f func(t T) InputOption) []InputOption {
	var options []InputOption
	for _, d := range data {
		options = append(options, f(d))
	}
	return options
}

func ProfessionalsToOptions(professionals []scheduling.Professional) []InputOption {
	return entityToOptions(professionals, func(p scheduling.Professional) InputOption {
		return InputOption{p.Name.String(), p.ID.String()}
	})
}

func ServicesToOptions(services []scheduling.Service) []InputOption {
	return entityToOptions(services, func(s scheduling.Service) InputOption {
		return InputOption{s.Name.String(), s.ID.String()}
	})
}

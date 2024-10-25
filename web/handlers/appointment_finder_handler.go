package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/components"
)

func NewAppointmentFinderHandler(f scheduling.AppointmentFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id := r.PathValue("id")

		o, err := f.Execute(id)

		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			http.Error(w, "Marcação não encontrada", http.StatusNotFound)
			return
		}

		if !errors.Is(nil, err) {
			http.Error(w, "Erro desconhecido, contacte o Administrador", http.StatusInternalServerError)
			return
		}

		components.RescheduleEventForm(o).Render(r.Context(), w)

	}
}

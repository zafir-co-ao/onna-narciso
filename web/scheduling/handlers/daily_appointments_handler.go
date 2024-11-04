package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
)

func HandleDailyAppointments(w http.ResponseWriter, r *http.Request) {
	pages.DailyAppointments().Render(r.Context(), w)
}

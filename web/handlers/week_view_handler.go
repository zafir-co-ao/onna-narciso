package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/components"
)

func HandleWeekView(w http.ResponseWriter, r *http.Request) {
	components.WeekView("2024-10-10", 6, 8, 22, []scheduling.Appointment{}).Render(r.Context(), w)
}

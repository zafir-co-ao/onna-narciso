package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/zafir-co-ao/onna-narciso/web/scheduling/components"
)

func HandleDailyAppointmentsCalendar() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		value := r.FormValue("operation")
		date, err := time.Parse("2006-01-02", r.FormValue("date"))
		state := components.CalendarState{
			HxGet:    r.FormValue("hx-get"),
			HxTarget: r.FormValue("hx-target"),
			HxSwap:   r.FormValue("hx-swap"),
			Date:     date.Format("2006-01-02"),
		}

		if !errors.Is(nil, err) {
			state.Date = date.Format("2006-01-02")
			components.Calendar(state).Render(r.Context(), w)
		}

		if value == "previous-month" {
			date = date.AddDate(0, -1, 0)
			state.Date = date.Format("2006-01-02")
		}

		if value == "next-month" {
			date = date.AddDate(0, 1, 0)
			state.Date = date.Format("2006-01-02")
		}

		components.Calendar(state).Render(r.Context(), w)
	}
}

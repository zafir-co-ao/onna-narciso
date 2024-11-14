package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/zafir-co-ao/onna-narciso/web/scheduling/components"
)

func HandleDailyAppointmentsCalendar() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		operation := r.FormValue("operation")
		state := components.CalendarState{
			HxGet:     r.FormValue("hx-get"),
			HxTarget:  r.FormValue("hx-target"),
			HxSwap:    r.FormValue("hx-swap"),
			HxTrigger: r.FormValue("hx-trigger"),
		}

		date, err := time.Parse("2006-01-02", r.FormValue("date"))
		if !errors.Is(nil, err) {
			state.Date = time.Now().Format("2006-01-02")
			components.Calendar(state).Render(r.Context(), w)
		}

		if operation == "previous-month" {
			date = date.AddDate(0, -1, 0)
			state.Date = date.Format("2006-01-02")
		}

		if operation == "next-month" {
			date = date.AddDate(0, 1, 0)
			state.Date = date.Format("2006-01-02")
		}

		components.Calendar(state).Render(r.Context(), w)
	}
}

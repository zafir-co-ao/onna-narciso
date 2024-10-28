package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/components"
)

func HandleDailyView(w http.ResponseWriter, r *http.Request) {
	components.DailyView().Render(r.Context(), w)
}

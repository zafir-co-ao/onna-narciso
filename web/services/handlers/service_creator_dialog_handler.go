package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/services/components"
)

func HandleCreateServiceDialog(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("hx-post")
	components.ServiceCreateDialog(url).Render(r.Context(), w)
}

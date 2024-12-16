package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/crm/components"
)

func HandleCreateCustomerDialog(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("hx-post")
	components.CustomerCreateDialog(url).Render(r.Context(), w)
}

package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/crm/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCreateCustomerDialog(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("hx-post")
	
	_http.SendOk(w)
	components.CustomerCreateDialog(url).Render(r.Context(), w)
}

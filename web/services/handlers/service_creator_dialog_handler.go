package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/services/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCreateServiceDialog(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("hx-post")

	_http.SendOk(w)
	components.ServiceCreateDialog(url).Render(r.Context(), w)
}

package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/auth/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleUserCreateDialog(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("hx-post")
	
	_http.SendOk(w)
	components.UserCreateDialog(url).Render(r.Context(), w)
}

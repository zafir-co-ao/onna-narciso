package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/auth/components"
)

func HandleUserCreateDialog(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("hx-post")
	components.UserCreateDialog(url).Render(r.Context(), w)
}

package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/auth/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleUserPasswordReset(w http.ResponseWriter, r *http.Request) {
	_http.SendOk(w)
	pages.ResetUserPassword().Render(r.Context(), w)
}

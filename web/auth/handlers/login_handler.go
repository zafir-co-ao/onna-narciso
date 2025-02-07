package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/auth/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	_http.SendOk(w)
	pages.Login().Render(r.Context(), w)
}

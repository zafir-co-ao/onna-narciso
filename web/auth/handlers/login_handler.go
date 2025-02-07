package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/auth/pages"
)

func HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	pages.Login().Render(r.Context(), w)
}

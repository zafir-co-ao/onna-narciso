package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/auth/pages"
)

func HandleUserPasswordReset(w http.ResponseWriter, r *http.Request) {
	pages.ResetUserPassword().Render(r.Context(), w)
}

package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/web/services/pages"
)

func HandleFindServices() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s := []services.ServiceOutput{}
		pages.ListService(s).Render(r.Context(), w)
	}
}

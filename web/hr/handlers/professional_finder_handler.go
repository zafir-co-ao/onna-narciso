package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/hr"
	"github.com/zafir-co-ao/onna-narciso/web/hr/pages"
)

func HandleFindProfessionals(u hr.ProfessionalFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		o, _ := u.FindAll()

		pages.ListProfessionals(o).Render(r.Context(), w)
	}
}

package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/web/services/components"
)

func HandleEditServiceDialog(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	url := r.FormValue("hx-post")

	s := services.ServiceOutput{
		ID:          id,
		Name:        "Pedicure",
		Duration:    90,
		Price:       "12344",
		Description: "Serviço",
	}

	components.ServiceEditDialog(url, s).Render(r.Context(), w)
}

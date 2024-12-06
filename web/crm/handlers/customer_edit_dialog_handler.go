package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/web/crm/components"
)

func HandleEditCustomerDialog() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("hx-put")
		_ = r.FormValue("id")

		o := crm.CustomerOutput{
			Name:        "Jonathan",
			Nif:         "008342442LA043",
			BirthDate:   "2001-04-17",
			Email:       "jonathan@gmail.com",
			PhoneNumber: "91234567",
		}

		components.HandleCustomerEditDialog(url, o).Render(r.Context(), w)
	}
}

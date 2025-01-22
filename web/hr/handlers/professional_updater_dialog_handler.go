package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/hr"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/web/hr/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleUpdateProfessionalDialog(pu hr.ProfessionalFinder, su services.ServiceFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		url := r.FormValue("hx-put")

		p, err := pu.FindByID(id)

		if errors.Is(err, hr.ErrProfessionalNotFound) {
			_http.SendBadRequest(w, "Profissional não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_services, err := su.FindAll()

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
		components.HandleUpdateProfessionalDialog(url, p, _services).Render(r.Context(), w)
	}
}

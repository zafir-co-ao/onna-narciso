package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/hr"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleUpdateProfessional(u hr.ProfessionalUpdater) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		i := hr.ProfessionalUpdaterInput{
			ID:          r.PathValue("id"),
			Name:        r.FormValue("name"),
			ServicesIDs: r.Form["serviceID"],
		}

		err := u.Update(i)

		if errors.Is(err, hr.ErrProfessionalNotFound) {
			_http.SendBadRequest(w, "Profissional não encontrado")
			return
		}

		if errors.Is(err, name.ErrEmptyName) {
			_http.SendBadRequest(w, "O nome do professional não pode ser vázio")
			return
		}

		if errors.Is(err, hr.ErrServiceNotFound) {
			_http.SendNotFound(w, "Serviço não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
	}
}

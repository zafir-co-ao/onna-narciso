package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/hr"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCreateProfessional(u hr.ProfessionalCreator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		i := hr.ProfessionalCreatorInput{
			Name:        r.FormValue("name"),
			ServicesIDs: r.Form["serviceID"],
		}

		_, err := u.Create(i)

		if errors.Is(err, name.ErrEmptyName) {
			_http.SendBadRequest(w, "O nome do profissional não pode ser vázio")
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

		w.Header().Set("X-Reload-Page", "ReloadPage")
		_http.SendCreated(w)
	}
}

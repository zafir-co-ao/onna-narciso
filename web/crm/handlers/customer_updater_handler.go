package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleUpdateCustomer(u crm.CustomerUpdater) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		i := crm.CustomerUpdaterInput{
			ID:          r.PathValue("id"),
			Name:        r.FormValue("name"),
			Nif:         r.FormValue("nif"),
			BirthDate:   r.FormValue("birth-date"),
			Email:       r.FormValue("email"),
			PhoneNumber: r.FormValue("phone-number"),
		}

		err := u.Update(i)

		if errors.Is(err, crm.ErrCustomerNotFound) {
			_http.SendNotFound(w, "Cliente não encontrado")
			return
		}

		if errors.Is(err, name.ErrEmptyName) {
			_http.SendBadRequest(w, "O nome do cliente não pode estar vazio")
			return
		}

		if errors.Is(err, crm.ErrNifAlreadyUsed) {
			_http.SendBadRequest(w, "O NIF fornecido já está sendo usado por um cliente diferente")
			return
		}

		if errors.Is(err, crm.ErrInvalidEmailFormat) {
			_http.SendBadRequest(w, "O e-mail fornecido é inválido")
			return
		}

		if errors.Is(err, date.ErrInvalidFormat) {
			_http.SendBadRequest(w, "A data de nascimento está no formato inválido")
			return
		}

		if errors.Is(err, crm.ErrEmailAlreadyUsed) {
			_http.SendBadRequest(w, "O e-mail fornecido já está sendo usado por um cliente diferente")
			return
		}

		if errors.Is(err, crm.ErrPhoneNumberAlreadyUsed) {
			_http.SendBadRequest(w, "O telefone fornecido já está sendo usado por um cliente diferente")
			return
		}

		if errors.Is(err, crm.ErrAgeNotAllowed) {
			_http.SendBadRequest(w, fmt.Sprintf("A idade mínima permitida é de %v anos", crm.MinimumAgeAllowed))
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
	}
}

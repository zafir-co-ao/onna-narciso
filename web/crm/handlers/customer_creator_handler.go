package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/email"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/nif"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/phone"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCreateCustomer(u crm.CustomerCreator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		i := crm.CustomerCreatorInput{
			Name:        r.Form.Get("name"),
			Nif:         r.Form.Get("nif"),
			BirthDate:   r.Form.Get("birth-date"),
			Email:       r.Form.Get("email"),
			PhoneNumber: r.Form.Get("phone-number"),
		}

		_, err := u.Create(i)

		if errors.Is(err, name.ErrEmptyName) {
			_http.SendBadRequest(w, "O nome do cliente não pode estar vazio")
			return
		}

		if errors.Is(err, nif.ErrEmptyNif) {
			_http.SendBadRequest(w, "O NIF do cliente não pode estar vazio")
			return
		}

		if errors.Is(err, crm.ErrNifAlreadyUsed) {
			_http.SendBadRequest(w, "O NIF fornecido já está sendo usado por um cliente diferente")
			return
		}

		if errors.Is(err, email.ErrInvalidFormat) {
			_http.SendBadRequest(w, "O e-mail fornecido é inválido")
			return
		}

		if errors.Is(err, phone.ErrEmptyPhoneNumber) {
			_http.SendBadRequest(w, "O telefone do cliente não pode estar vazio")
			return
		}

		if errors.Is(err, date.ErrInvalidFormat) {
			_http.SendBadRequest(w, "A data de nascimento está no formato inválido")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendCreated(w)
	}
}
package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/services/price"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCreateService(u services.ServiceCreator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		d, err := strconv.Atoi(r.FormValue("duration"))
		if err != nil {
			_http.SendBadRequest(w, "A duração do serviço está no formato inválido")
			return
		}

		i := services.ServiceCreatorInput{
			Name:        r.Form.Get("name"),
			Price:       r.Form.Get("price"),
			Description: r.Form.Get("description"),
			Duration:    d,
		}

		_, err = u.Create(i)

		if errors.Is(name.ErrInvalidName, err) {
			_http.SendBadRequest(w, "O nome do serviço não pode estar vazio")
			return
		}

		if errors.Is(price.ErrInvalidPrice, err) {
			_http.SendBadRequest(w, "O preço do serviço está no formato inválido")
			return
		}

		if errors.Is(duration.ErrInvalidDuration, err) {
			_http.SendBadRequest(w, "A duração do serviço não deve ser inferior a zero")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendCreated(w)
	}
}
package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zafir-co-ao/onna-narciso/internal/services"
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
			Name:        r.FormValue("name"),
			Price:       r.FormValue("price"),
			Description: r.FormValue("description"),
			Discount:    r.FormValue("discount"),
			Duration:    d,
		}

		_, err = u.Create(i)

		if errors.Is(err, name.ErrEmptyName) {
			_http.SendBadRequest(w, "O nome do serviço não pode estar vazio")
			return
		}

		if errors.Is(err, services.ErrInvalidPrice) {
			_http.SendBadRequest(w, "O preço do serviço está no formato inválido")
			return
		}

		if errors.Is(err, services.ErrDiscountNotAllowed) {
			_http.SendBadRequest(w, "O valor de desconto não é válido")
			return
		}

		if errors.Is(err, duration.ErrInvalidDuration) {
			_http.SendBadRequest(w, "A duração do serviço não deve ser inferior a zero")
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

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

func HandleEditService(u services.ServiceEditor) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		d, err := strconv.Atoi(r.FormValue("duration"))
		if err != nil {
			_http.SendBadRequest(w, "A duração do serviço está no formato inválido")
			return
		}

		i := services.ServiceEditorInput{
			ID:          id,
			Name:        r.Form.Get("name"),
			Price:       r.Form.Get("price"),
			Description: r.Form.Get("description"),
			Duration:    d,
		}

		err = u.Edit(i)

		if errors.Is(err, name.ErrEmptyName) {
			_http.SendBadRequest(w, "O nome do serviço não pode estar vazio")
			return
		}

		if errors.Is(err, services.ErrInvalidPrice) {
			_http.SendBadRequest(w, "O preço do serviço está no formato inválido")
			return
		}

		if errors.Is(err, duration.ErrInvalidDuration) {
			_http.SendBadRequest(w, "A duração do serviço não deve ser inferior a zero")
			return
		}

		if errors.Is(err, services.ErrServiceNotFound) {
			_http.SendBadRequest(w, "Serviço não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		w.Header().Set("X-Reload-Page", "ReloadPage")
		_http.SendOk(w)
	}
}

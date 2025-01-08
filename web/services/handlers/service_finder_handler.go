package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/web/services/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleFindServices(u services.ServiceFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		o, err := u.FindAll()

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		pages.ListServices(o).Render(r.Context(), w)
		_http.SendOk(w)
	}
}

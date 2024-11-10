package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/session"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCloseSession(c session.Closer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		input := session.CloserInput{SessionID: r.PathValue("id")}

		err := c.Close(input)

		if errors.Is(session.ErrSessionClosed, err) {
			_http.SendBadRequest(w, "A sessão já foi finalizada")
			return
		}

		if errors.Is(session.ErrSessionNotFound, err) {
			_http.SendNotFound(w, "A sessão não foi encontrada")
			return
		}

		if errors.Is(session.ErrServiceNotFound, err) {
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

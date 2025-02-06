package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleGetAuthenticatedUser(w http.ResponseWriter, r *http.Request, uf auth.UserFinder) (auth.UserOutput, bool) {
	cookie, err := r.Cookie("userID")

	if !errors.Is(nil, err) {
		_http.SendUnauthorized(w)
		return auth.UserOutput{}, false
	}

	uid := cookie.Value

	au, err := uf.FindByID(uid)

	if errors.Is(err, auth.ErrUserNotFound) {
		_http.SendNotFound(w, "Utilizador não encontrado")
		return auth.UserOutput{}, false
	}

	if !errors.Is(nil, err) {
		_http.SendServerError(w)
		return auth.UserOutput{}, false
	}

	return au, true
}

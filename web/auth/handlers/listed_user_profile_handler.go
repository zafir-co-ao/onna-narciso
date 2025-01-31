package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/web/auth/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleListedUserProfilePage(uf auth.UserFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		undefinedID := "${id}"

		cookie, _ := r.Cookie("profileID")

		if id == undefinedID {
			id = cookie.Value
		}

		o, err := uf.FindByID(id)

		if errors.Is(err, auth.ErrUserNotFound) {
			_http.SendNotFound(w, "Utilizador não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		au, ok := GetAuthenticatedUser(w, r, uf)
		if !ok {
			return
		}

		_http.SendOk(w)
		pages.ListedUserProfile(o, au).Render(r.Context(), w)
	}
}

func GetAuthenticatedUser(w http.ResponseWriter, r *http.Request, uf auth.UserFinder) (auth.UserOutput, bool) {
	cookie, _ := r.Cookie("userID")

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

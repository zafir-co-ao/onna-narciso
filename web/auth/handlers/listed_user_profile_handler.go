package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/web/auth/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleListedUserProfilePage(u auth.UserFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		undefinedID := "${id}"

		cookie := &http.Cookie{
			Name:     "profileID",
			Value:    id,
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		}

		http.SetCookie(w, cookie)

		cookie, _ = r.Cookie("ProfileID")

		if id == undefinedID {
			id = cookie.Value
			fmt.Println("Cookie: ", id)
		}

		fmt.Println("ID: ", id)

		o, err := u.FindByID(id)

		if errors.Is(err, auth.ErrUserNotFound) {
			_http.SendNotFound(w, "Utilizador não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
		pages.ListedUserProfile(o).Render(r.Context(), w)
	}
}

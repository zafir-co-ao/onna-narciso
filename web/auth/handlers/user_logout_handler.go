package handlers

import (
	"net/http"
	"time"

	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleLogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:    "userID",
		Value:   "",
		Expires: time.Now().AddDate(0, 0, -1),
		MaxAge:  -1,
	}

	http.SetCookie(w, cookie)

	w.Header().Set("HX-Redirect", "/auth/login")
	_http.SendOk(w)
}

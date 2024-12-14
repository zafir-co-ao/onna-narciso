package http

import "net/http"

func SendReponse(w http.ResponseWriter, msg string, status int) {
	http.Error(w, msg, status)
}

func SendNotFound(w http.ResponseWriter, msg string) {
	SendReponse(w, msg, http.StatusNotFound)
}

func SendBadRequest(w http.ResponseWriter, msg string) {
	SendReponse(w, msg, http.StatusBadRequest)
}

func SendUnauthorized(w http.ResponseWriter) {
	SendReponse(w, "Crêdencias inválidas", http.StatusUnauthorized)
}

func SendMethodNotAllowed(w http.ResponseWriter) {
	SendReponse(w, "method not allowed", http.StatusMethodNotAllowed)
}

func SendServerError(w http.ResponseWriter) {
	SendReponse(w, "Erro desconhecido, contacte o Administrador", http.StatusInternalServerError)
}

func SendOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func SendCreated(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
}

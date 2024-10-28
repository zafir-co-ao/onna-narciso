package handlers

import "net/http"

func sendReponse(w http.ResponseWriter, msg string, status int) {
	http.Error(w, msg, status)
}

func sendNotFound(w http.ResponseWriter, msg string) {
	sendReponse(w, msg, http.StatusNotFound)
}

func sendBadRequest(w http.ResponseWriter, msg string) {
	sendReponse(w, msg, http.StatusBadRequest)
}

func sendMethodNotAllowed(w http.ResponseWriter) {
	sendReponse(w, "method not allowed", http.StatusMethodNotAllowed)
}

func sendServerError(w http.ResponseWriter) {
	sendReponse(w, "Erro desconhecido, contacte o Administrador", http.StatusInternalServerError)
}

func sendOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func sendCreated(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
}

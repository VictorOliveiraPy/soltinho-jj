package errors

import "net/http"

func HandleConflictError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusConflict)
	w.Write([]byte(err.Error()))
}

func HandleUnauthorizedError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(err.Error()))
}

func HandleNotFoundError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(err.Error()))
}

func HandleInternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func HandleBadRequestError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

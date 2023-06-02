package errors

import "net/http"

func HandleConflictError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusConflict)
	w.Write([]byte(err.Error()))
}

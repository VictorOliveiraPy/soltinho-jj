package handlers

import (
	"github.com/VictorOliveiraPy/internal/infra/logger"
	"net/http"
)

func BadRequestHandler(w http.ResponseWriter, err error) {
	logger.Error("Erro de aplicação - Bad Request", err)
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func NotFoundHandler(w http.ResponseWriter, err error) {
	logger.Error("Erro de aplicação - Not Found", err)
	http.Error(w, err.Error(), http.StatusNotFound)
}

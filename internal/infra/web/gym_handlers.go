package web

import (
	"encoding/json"
	"net/http"

	"github.com/VictorOliveiraPy/internal/errors"
	"github.com/VictorOliveiraPy/internal/service"
	"github.com/VictorOliveiraPy/internal/usecase"
	"github.com/go-chi/jwtauth"
)

type WebGymHandler struct {
	GymService *service.GymService
	TokenAuth  *jwtauth.JWTAuth
}

func NewWebGymHandler(gymService *service.GymService, tokenAuth *jwtauth.JWTAuth) *WebGymHandler {
	return &WebGymHandler{
		GymService: gymService,
		TokenAuth:  tokenAuth,
	}
}

func (h *WebGymHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.GymInput

	_, claims, _ := jwtauth.FromContext(r.Context())
	user_id := claims["user"].(string)

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		errors.HandleBadRequestError(w, err)
		return
	}

	err = h.GymService.CreateGym(r.Context(), dto, user_id)
	if err != nil {
		if e, ok := err.(errors.UnauthorizedError); ok {
			errors.HandleUnauthorizedError(w, e)
			return
		}
		if e, ok := err.(errors.GymNameAlreadyExistsError); ok {
			errors.HandleConflictError(w, e)
		} else {
			errors.HandleInternalServerError(w, err)
		}
	}

	w.WriteHeader(http.StatusCreated)

}

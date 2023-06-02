package web

import (
	"encoding/json"

	"net/http"

	"github.com/VictorOliveiraPy/internal/errors"
	"github.com/VictorOliveiraPy/internal/service"
	"github.com/go-chi/jwtauth"

	"github.com/VictorOliveiraPy/internal/usecase"
)

type WebUserHandler struct {
	UserService *service.UserService
	UseCase     *usecase.GetTokenUseCase
}

func NewWebUserHandler(userService *service.UserService, usecase *usecase.GetTokenUseCase) *WebUserHandler {
	return &WebUserHandler{
		UserService: userService,
		UseCase:     usecase,
	}
}

// GetJWT godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body usecase.GetJWTInput  true  "user credentials"
// @Failure      404  {object}  Error
// @Failure      403 {object}  Error
// @Router       /users/generate_token [post]
func (h *WebUserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExperesIn").(int)
	var dto usecase.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessToken, err := h.UseCase.GetUserToken(jwt, jwtExpiresIn, dto)
	if e, ok := err.(errors.EmailNotFound); ok {
		errors.HandleNotFoundError(w, e)
		return
	}
	if e, ok := err.(errors.PasswordInvalid); ok {
		errors.HandleUnauthorizedError(w, e)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}

func (h *WebUserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.UserInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		errors.HandleBadRequestError(w, err)
		return
	}

	err = h.UserService.CreateUser(r.Context(), dto)
	if err != nil {
		if e, ok := err.(errors.EmailAlreadyExistsError); ok {
			errors.HandleConflictError(w, e)
			return
		} else {
			errors.HandleInternalServerError(w, err)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

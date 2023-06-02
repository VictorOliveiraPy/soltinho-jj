package web

import (
	"encoding/json"
	"net/http"

	"github.com/VictorOliveiraPy/internal/errors"
	"github.com/VictorOliveiraPy/internal/service"
	"github.com/VictorOliveiraPy/internal/usecase"
	"github.com/go-chi/jwtauth"
)


type WebStudentHandler struct {
	StudentService *service.StudentService
	TokenAuth  *jwtauth.JWTAuth
}

func NewWebStudentHandler(studentService *service.StudentService, tokenAuth *jwtauth.JWTAuth) *WebStudentHandler {
	return &WebStudentHandler{
		StudentService: studentService,
		TokenAuth:  tokenAuth,
	}
}

func (h *WebStudentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.StudentInput

	_, claims, _ := jwtauth.FromContext(r.Context())
	user_id := claims["user"].(string)

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		errors.HandleBadRequestError(w, err)
		return
	}

	err = h.StudentService.CreateStudent(r.Context(), dto, user_id)
	if err != nil {
		if e, ok := err.(errors.UnauthorizedError); ok {
			errors.HandleUnauthorizedError(w, e)
			return
		} else {
			errors.HandleInternalServerError(w, err)
		}
	}

	w.WriteHeader(http.StatusCreated)

}

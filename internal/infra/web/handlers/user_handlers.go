package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/VictorOliveiraPy/internal/dto"
	"github.com/VictorOliveiraPy/internal/entity"

	db "github.com/VictorOliveiraPy/internal/infra/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	userDB *sql.DB
	*db.Queries
	Jwt *jwtauth.JWTAuth
}

type Error struct {
	Message string `json:"message"`
}

func NewUserHandler(dbConn *sql.DB) *UserHandler {
	return &UserHandler{
		userDB:  dbConn,
		Queries: db.New(dbConn),
	}
}


func (handler *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExperesIn").(int)

	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := handler.Queries.GetUserByEmail(ctx, user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID,
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}


func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var request dto.User

	json.NewDecoder(r.Body).Decode(&request)

	email, _ := h.Queries.GetUserByEmail(ctx, request.Email)

	if email.Email != ""{
        w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "O email fornecido já está em uso.")

        return
    }

	user, err := entity.NewUser(request.Username, request.Password, request.Email, request.RoleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	err = h.Queries.CreateUser(ctx, db.CreateUserParams{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.UserName,
		Password: user.Password,
		RoleID:   user.RoleID,
		Active:   user.Active,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (h *UserHandler) GetUserFullProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&id)

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	row, err := h.Queries.GetUserCompleteProfile(context.Background(), id)
	if err!= nil {
		w.WriteHeader(http.StatusNotFound)
        err := Error{Message: err.Error()}
        json.NewEncoder(w).Encode(err)
        return
	}

	dtoResponse := dto.UserCompleteProfile{
		ID:           row.ID,
		Username:     row.Username,
		Email:        row.Email,
		RoleID:       row.RoleID,
		Active:       row.Active,
		GymID:        row.GymID,
		GymName:        sql.NullString{String: row.GymName.String, Valid: row.GymName.Valid},
		TeamName:       sql.NullString{String: row.TeamName.String, Valid: row.TeamName.Valid},
		GymActive:      sql.NullBool{Bool: row.GymActive.Bool, Valid: row.GymActive.Valid},
		StudentID:      sql.NullString{String: row.StudentID.String, Valid: row.StudentID.Valid},
		Graduation:     sql.NullString{String: row.Graduation.String, Valid: row.Graduation.Valid},
		StudentActive:  sql.NullBool{Bool: row.StudentActive.Bool, Valid: row.StudentActive.Valid},
		TrainingTime:   sql.NullString{String: row.TrainingTime.String, Valid: row.TrainingTime.Valid},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dtoResponse)

}

package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/VictorOliveiraPy/internal/dto"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/infra/logger"
	"go.uber.org/zap"

	db "github.com/VictorOliveiraPy/internal/infra/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}



func (handler *EntityHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExperesIn").(int)
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		BadRequestHandler(w, err)
		return
	}

	u, err := handler.Queries.GetUserByEmail(context.Background(), user.Email)
	if err != nil {
		NotFoundHandler(w, err)
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

func (h *EntityHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("New user creation initiated",
		zap.String("route", "/create-user"),
	)
	ctx := context.Background()
	var request dto.User

	json.NewDecoder(r.Body).Decode(&request)

	email, _ := h.Queries.GetUserByEmail(ctx, request.Email)

	if email.Email != "" {
		err := fmt.Errorf(strings.TrimRight("O email fornecido já está em uso.", "\n.,;:!?"))
		logger.Error("Erro durante a criação do usuário", err,
			zap.String("email", email.Email),
		)
		w.WriteHeader(http.StatusConflict)
		return
	}

	user, err := entity.NewUser(request.Username, request.Password, request.Email, request.RoleID)
	if err != nil {
		logger.Error("Erro ao criar objetos de um novo usuário", err)
		BadRequestHandler(w, err)
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
		logger.Error("Erro ao criar um novo usuário", err)
		BadRequestHandler(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	logger.Info("New user created",
		zap.String("route", "/create-user"),
		zap.String("user_id", user.ID),
	)

}

func (h *EntityHandler) GetUserFullProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&id)

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	row, err := h.Queries.GetUserCompleteProfile(context.Background(), id)
	if err != nil {
		err := fmt.Errorf(strings.TrimRight("ID fornecido Nao existe.", "\n.,;:!?"))
		logger.Error("ID fornecido Nao existe.", err,
			zap.String("ID", id),
		)
		NotFoundHandler(w, err)
		return
	}

	dtoResponse := dto.UserCompleteProfile{
		ID:            row.ID,
		Username:      row.Username,
		Email:         row.Email,
		RoleID:        row.RoleID,
		Active:        row.Active,
		GymID:         row.GymID,
		GymName:       sql.NullString{String: row.GymName.String, Valid: row.GymName.Valid},
		TeamName:      sql.NullString{String: row.TeamName.String, Valid: row.TeamName.Valid},
		GymActive:     sql.NullBool{Bool: row.GymActive.Bool, Valid: row.GymActive.Valid},
		StudentID:     sql.NullString{String: row.StudentID.String, Valid: row.StudentID.Valid},
		Graduation:    sql.NullString{String: row.Graduation.String, Valid: row.Graduation.Valid},
		StudentActive: sql.NullBool{Bool: row.StudentActive.Bool, Valid: row.StudentActive.Valid},
		TrainingTime:  sql.NullString{String: row.TrainingTime.String, Valid: row.TrainingTime.Valid},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dtoResponse)

}

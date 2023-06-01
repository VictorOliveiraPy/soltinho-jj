// package handlers

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/VictorOliveiraPy/internal/dto"
// 	"github.com/VictorOliveiraPy/internal/entity"
// 	"github.com/go-chi/chi/v5"
// 	"go.uber.org/zap"

// 	db "github.com/VictorOliveiraPy/internal/infra/database"
// 	"github.com/VictorOliveiraPy/internal/infra/logger"
// )

// func (h *EntityHandler) CreateGym(w http.ResponseWriter, r *http.Request) {
// 	var gym dto.RequestGym

// 	err := json.NewDecoder(r.Body).Decode(&gym)
// 	if err != nil {
// 		BadRequestHandler(w, err)
// 		return
// 	}

// 	role, err := h.Queries.GetUserRoleName(context.Background(), gym.UserID)

// 	if err != nil {
// 		NotFoundHandler(w, err)
// 		return
// 	}

// 	if !entity.IsAdminOrInstructor(role) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		err := Error{Message: err.Error()}
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}

// 	gym_name, _ := h.Queries.GetByGymName(context.Background(), gym.GymName)
// 	if gym_name.GymName != "" {
// 		err := fmt.Errorf(strings.TrimRight("O Nome da Academia fornecido já está em uso.", "\n.,;:!?"))
// 		logger.Error("Erro durante a criação da academia", err,
// 			zap.String("GymName", gym.GymName),
// 		)
// 		w.WriteHeader(http.StatusConflict)
// 		return
// 	}

// 	gymNew, err := entity.NewGym(gym.UserID, gym.GymName, gym.TeamName)

// 	if err != nil {
// 		BadRequestHandler(w, err)
// 		return
// 	}

// 	err = h.Queries.CreateGym(context.Background(), db.CreateGymParams{
// 		ID:       gymNew.ID,
// 		UserID:   gymNew.UserID,
// 		GymName:  gymNew.GymName,
// 		TeamName: gymNew.TeamName,
// 		Active:   gymNew.Active,
// 	})

// 	if err != nil {
// 		BadRequestHandler(w, err)
// 		return
// 	}

// }

// func (h *EntityHandler) GetByGym(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	json.NewDecoder(r.Body).Decode(&id)

// 	if id == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}

// 	gym, err := h.Queries.GetGymByID(context.Background(), id)
// 	if err != nil {
// 		err := fmt.Errorf(strings.TrimRight("ID fornecido Nao existe.", "\n.,;:!?"))
// 		logger.Error("ID fornecido Nao existe.", err,
// 			zap.String("ID", id),
// 		)
// 		NotFoundHandler(w, err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(gym)

// }

// func (h *EntityHandler) GetAllGyms(w http.ResponseWriter, r *http.Request) {
// 	gym, err := h.Queries.GetAllGyms(context.Background())
// 	if err != nil {
// 		http.Error(w, "Erro ao consultar as academias", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(gym)
// }

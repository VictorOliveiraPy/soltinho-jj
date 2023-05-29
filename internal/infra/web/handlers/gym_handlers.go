package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/VictorOliveiraPy/internal/dto"
	"github.com/VictorOliveiraPy/internal/entity"

	db "github.com/VictorOliveiraPy/internal/infra/database"
)



func (h *EntityHandler) CreateGym(w http.ResponseWriter, r *http.Request) {
	var gym dto.RequestGym

	err := json.NewDecoder(r.Body).Decode(&gym)
	if err != nil {
		BadRequestHandler(w, err)
		return
	}

	role, err := h.Queries.GetUserRoleName(context.Background(), gym.UserID)

	if err != nil {
		NotFoundHandler(w, err)
		return
	}

	if !entity.IsAdminOrInstructor(role) {
		w.WriteHeader(http.StatusUnauthorized)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	gymNew, err := entity.NewGym(gym.UserID, gym.GymName, gym.TeamName)

	if err != nil {
		BadRequestHandler(w, err)
		return
	}

	err = h.Queries.CreateGym(context.Background(), db.CreateGymParams{
		ID:       gymNew.ID,
		UserID:   gymNew.UserID,
		GymName:  gymNew.GymName,
		TeamName: gymNew.TeamName,
	})

	if err != nil {
		BadRequestHandler(w, err)
		return
	}

}

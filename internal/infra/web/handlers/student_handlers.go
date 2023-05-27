package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/VictorOliveiraPy/internal/dto"
	"github.com/VictorOliveiraPy/internal/entity"
	db "github.com/VictorOliveiraPy/internal/infra/database"
)

type StudentHandler struct {
	studentDB *sql.DB
	*db.Queries
}

func NewStudentHandler(dbConn *sql.DB) *StudentHandler {
	return &StudentHandler{
		studentDB: dbConn,
		Queries:   db.New(dbConn),
	}
}

func (h *StudentHandler) Createstudent(w http.ResponseWriter, r *http.Request) {
	var request dto.Student

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	println("ate aqui cheguei1")
	_ , err = h.Queries.GetGymByID(context.Background(), request.GymID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	role, err := h.Queries.GetUserRoleName(context.Background(), request.UserID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if !entity.IsAdminOrInstructor(role) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}

	new_student, err := entity.NewStudent(request.GymID, request.GymName, request.Graduation, request.TrainingTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	err = h.Queries.CreateStudent(context.Background(), db.CreateStudentParams{
		ID:           new_student.ID,
		GymID:        new_student.GymID,
		Name:         new_student.GymName,
		Graduation:   new_student.Graduation,
		TrainingTime: new_student.TrainingTime,
		Active: new_student.Active,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

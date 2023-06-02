package entity

import (
	"github.com/google/uuid"
)

type Student struct {
	ID           string `json:"id"`
	GymID        string `json:"gym_id"`
	Name         string `json:"name"`
	Graduation   string `json:"graduation"`
	Active       bool   `json:"active"`
	TrainingTime string `json:"training_time"`
}

func NewStudent(gymID string, Name string, graduation string, trainingTime string) (*Student, error) {
	return &Student{
		ID:           uuid.New().String(),
		GymID:        gymID,
		Name:      	  Name,
		Graduation:   graduation,
		Active:       true,
		TrainingTime: trainingTime,
	}, nil
}

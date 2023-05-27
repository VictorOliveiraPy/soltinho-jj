package entity

import (
	"github.com/google/uuid"
)

type Student struct {
	ID           string `json:"id"`
	GymID        string `json:"gym_id"`
	GymName      string `json:"gym_name"`
	Graduation   string `json:"graduation"`
	Active       bool   `json:"active"`
	TrainingTime string `json:"training_time"`
}

func NewStudent(gymID string, gym_name string, graduation string, trainingTime string) (*Student, error) {
	return &Student{
		ID:           uuid.New().String(),
		GymID:        gymID,
		GymName:      gym_name,
		Graduation:   graduation,
		Active:       true,
		TrainingTime: trainingTime,
	}, nil
}

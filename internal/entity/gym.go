package entity

import "github.com/google/uuid"

type Gym struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	GymName  string `json:"gym_name"`
	TeamName string `json:"team_name"`
	Active   bool   `json:"active"`
}

func NewGym(user_id string, gym_name string, team_name string) (*Gym, error) {
	return &Gym{
		ID:       uuid.New().String(),
		UserID:   user_id,
		GymName:  gym_name,
		TeamName: team_name,
		Active:   true,
	}, nil
}

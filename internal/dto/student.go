package dto

type Student struct {
	ID           string `json:"id"`
	GymID        string `json:"gym_id"`
	GymName      string `json:"gym_name"`
	Graduation   string `json:"graduation"`
	Active       bool   `json:"active"`
	TrainingTime string `json:"training_time"`
	UserID       string `json:"user_id"`
}

package dto

type RequestGym struct {
	UserID   string `json:"user_id"`
	GymName  string `json:"gym_name"`
	TeamName string `json:"team_name"`
}

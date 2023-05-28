package dto

import "database/sql"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
	RoleID   string `json:"role_id"`
	Active   bool   `json:"active"`
}

type GetJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"-"`
}

type GetJWTOutput struct {
	AccessToken string `json:"acess_token"`
}


type UserCompleteProfile struct {
	ID             string         `json:"id"`
	Username       string         `json:"username"`
	Email          string         `json:"email"`
	RoleID         string         `json:"role_id"`
	Active         bool           `json:"active"`
	GymID          sql.NullString `json:"gym_id"`
	GymName        sql.NullString `json:"gym_name"`
	TeamName       sql.NullString `json:"team_name"`
	GymActive      sql.NullBool   `json:"gym_active"`
	StudentID      sql.NullString `json:"student_id"`
	Graduation     sql.NullString `json:"graduation"`
	StudentActive  sql.NullBool   `json:"student_active"`
	TrainingTime   sql.NullString `json:"training_time"`
}

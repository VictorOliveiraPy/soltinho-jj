package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
	RoleID   string `json:"role_id"`
	Active   bool   `json:"active"`
}

func NewUser(user_name string, password string, email string, role_id string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       uuid.New().String(),
		UserName: user_name,
		Password: string(hash),
		Email:    email,
		RoleID:   role_id,
		Active:   true,
	}, nil
}

func IsAdminOrInstructor(name_role string) bool {
	return name_role == "admin" || name_role == "instructor"
}

package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	AcademyName    string `json:"academy_name"`
	InstructorBelt string `json:"instructor_belt"`
	Password       string `json:"-"`
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func NewUser(name, email, phone, academy_name, instructor_belt, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:             uuid.New().String(),
		Name:           name,
		Email:          email,
		Phone:          phone,
		AcademyName:    academy_name,
		InstructorBelt: instructor_belt,
		Password:       string(hash),
	}, nil
}

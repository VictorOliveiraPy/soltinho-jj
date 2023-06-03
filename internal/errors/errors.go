package errors

import (
	"fmt"
)

type EmailNotFound struct {
	Email string
}

func (e EmailNotFound) Error() string {
	return fmt.Sprintf("o email '%s' não foi encontrado", e.Email)
}

type PasswordInvalid struct {
	Password string
}

func (e PasswordInvalid) Error() string {
	return fmt.Sprintf("a senha '%s' é inválida", e.Password)
}

type EmailAlreadyExistsError struct {
	Email string
}

func (e EmailAlreadyExistsError) Error() string {
	return fmt.Sprintf("o email '%s' já está em uso", e.Email)
}

type UnauthorizedError struct{}

func (e UnauthorizedError) Error() string {
	return "Usuário não autorizado"
}

func ErrUnauthorized() error {
	return UnauthorizedError{}
}

type GymNameAlreadyExistsError struct {
	GymName string
}

func (e GymNameAlreadyExistsError) Error() string {
	return fmt.Sprintf("O nome da academia '%s' já existe", e.GymName)
}

type GymNotFoundError struct {
	GymID string
}

func (e GymNotFoundError) Error() string {
	return fmt.Sprintf("Gym not found with ID: %s", e.GymID)
}

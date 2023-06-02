package errors

import "fmt"

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


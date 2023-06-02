package usecase

import (
	"github.com/VictorOliveiraPy/internal/entity"
)

type UserInputDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleID   string `json:"role_id"`
	Active   bool   `json:"active"`
}

type UserOutputDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
	RoleID   string `json:"role_id"`
	Active   bool   `json:"active"`
}

type CreateUserUseCase struct {
	UserRepository entity.UserRepositoryInterface
}

func NewCreateUserUseCase(
	UserRepository entity.UserRepositoryInterface,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: UserRepository,
	}
}

func (c *CreateUserUseCase) Execute(input UserInputDTO) error {
	user := entity.User{
		ID:       input.ID,
		UserName: input.Username,
		Password: input.Password,
		Email:    input.Email,
		RoleID:   input.RoleID,
		Active:   input.Active,
	}

	err := c.UserRepository.Create(&user)
	if err != nil {
		return err
	}

	return nil

}

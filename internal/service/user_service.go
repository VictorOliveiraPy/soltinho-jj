package service

import (
	"context"
	"fmt"

	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/usecase"
)


type UserService struct {
	CreateUserUseCase usecase.CreateUserUseCase
}

func NewUserService(createUserUseCase usecase.CreateUserUseCase) *UserService {
	return &UserService{
		CreateUserUseCase: createUserUseCase,
	}
}

func (s *UserService) CreateUser(ctx context.Context,dto usecase.UserInputDTO) error {

	println("cheguei aqui")
	user, err := entity.NewUser(dto.Username, dto.Password, dto.Email, dto.RoleID)
	if err != nil {
		return fmt.Errorf("erro ao criar novo usuário: %w", err)
	}

	err = s.CreateUserUseCase.Execute(usecase.UserInputDTO{
		ID:       user.ID,
		Username: user.UserName,
		Password: user.Password,
		Email:    user.Email,
		RoleID:   user.RoleID,
		Active:   user.Active,
	})

	if err != nil {
		return fmt.Errorf("erro ao criar um novo usuário: %w", err)
	}
	return nil
}
package service

import (
	"context"
	"fmt"

	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/usecase"
)


type EmailAlreadyExistsError struct {
	Email string
}

func (e EmailAlreadyExistsError) Error() string {
	return fmt.Sprintf("o email '%s' já está em uso", e.Email)
}

type UserService struct {
	CreateUserUseCase usecase.CreateUserUseCase
	userRepository entity.UserRepositoryInterface
}

func NewUserService(createUserUseCase usecase.CreateUserUseCase, userRepository entity.UserRepositoryInterface) *UserService {
	return &UserService{
		CreateUserUseCase: createUserUseCase,
		userRepository: userRepository,
	}
}

func (s *UserService) checkEmailExists(email string) error {
	existingUser, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return EmailAlreadyExistsError{Email: email}
	}
	return nil
}


func (s *UserService) CreateUser(ctx context.Context,dto usecase.UserInputDTO) error {
	if err := s.checkEmailExists(dto.Email); err != nil {
		println(err.Error())
		return err
	}

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
package service

import (
	"context"
	"fmt"

	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/errors"
	"github.com/VictorOliveiraPy/internal/usecase"
)

type GymService struct {
	CreateGymUseCase usecase.CreateGymUseCase
	gymRepository    entity.GymRepositoryInterface
	userRepository   entity.UserRepositoryInterface
}

func NewGymService(createGymUseCase usecase.CreateGymUseCase, gymRepository entity.GymRepositoryInterface, userRepository entity.UserRepositoryInterface) *GymService {
	return &GymService{
		CreateGymUseCase: createGymUseCase,
		gymRepository:    gymRepository,
		userRepository:   userRepository,
	}
}

func (s *GymService) CheckIfGymNameExists(gymName string) error {
	gym, err := s.gymRepository.FindByName(gymName)
	if err != nil {
		return err
	}

	if gym != nil && gym.GymName != "" {
		return errors.GymNameAlreadyExistsError{GymName: gymName}
	}

	return nil
}

func (s *GymService) CheckUserRole(userID string) error {
	role, err := s.userRepository.FindById(userID)
	if err != nil {
		return err
	}
	var user entity.User

	if !user.IsAuthorizedRole(role.RoleID) {
		return errors.ErrUnauthorized()
	}

	return nil
}

func (s *GymService) CreateGym(ctx context.Context, dto usecase.GymInput, userID string) error {
	err := s.CheckUserRole(userID)
	if err != nil {
		return err
	}

	err = s.CheckIfGymNameExists(dto.GymName)
	if err != nil {
		return err
	}

	gym, err := entity.NewGym(userID, dto.GymName, dto.TeamName)
	if err != nil {
		return fmt.Errorf("erro ao criar novo usuário: %w", err)
	}

	err = s.CreateGymUseCase.Execute(usecase.GymInput{
		ID:       gym.ID,
		UserID:   gym.UserID,
		GymName:  gym.GymName,
		TeamName: gym.TeamName,
		Active:   gym.Active,
	})

	if err != nil {
		return fmt.Errorf("erro ao criar uma nova Academia : %w", err)
	}
	return nil
}

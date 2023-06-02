package service

import (
	"context"
	"fmt"

	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/errors"
	"github.com/VictorOliveiraPy/internal/usecase"
)

type StudentService struct {
	CreateStudentUseCase usecase.CreateStudentUseCase
	StudentRepository    entity.StudentRepositoryInterface
	userRepository   entity.UserRepositoryInterface
}

func NewStudentService(createStudentUseCase usecase.CreateStudentUseCase, StudentRepository entity.StudentRepositoryInterface, userRepository entity.UserRepositoryInterface) *StudentService {
	return &StudentService{
		CreateStudentUseCase: createStudentUseCase,
		StudentRepository:    StudentRepository,
		userRepository:   userRepository,
	}
}


func (s *StudentService) CheckUserRole(userID string) error {
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

func (s *StudentService) CreateStudent(ctx context.Context, dto usecase.StudentInput, userID string) error {
	err := s.CheckUserRole(userID)
	if err != nil {
		return err
	}

	student, err := entity.NewStudent(dto.GymID, dto.Name, dto.Graduation, dto.TrainingTime)
	if err != nil {
		return fmt.Errorf("erro ao criar novo usuário: %w", err)
	}


	err = s.CreateStudentUseCase.Execute(usecase.StudentInput{
		ID: student.ID,
		GymID: student.GymID,
		Name: student.Name,
		Graduation: student.Graduation,
		Active: student.Active,
		TrainingTime: student.TrainingTime,
	})

	if err != nil {
		return fmt.Errorf("erro ao criar um aluno : %w", err)
	}
	return nil
}

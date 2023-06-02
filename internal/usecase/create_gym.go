package usecase

import "github.com/VictorOliveiraPy/internal/entity"

type CreateGymUseCase struct {
	GymRepository entity.GymRepositoryInterface
}

func NewCreaGymUseCase(
	GymRepository entity.GymRepositoryInterface,
) *CreateGymUseCase {
	return &CreateGymUseCase{
		GymRepository: GymRepository,
	}
}

type GymInput struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	GymName  string `json:"gym_name"`
	TeamName string `json:"team_name"`
	Active   bool   `json:"active"`
}

func (c *CreateGymUseCase) Execute(input GymInput) error {
	gym := entity.Gym{
		ID:       input.ID,
		UserID:   input.UserID,
		GymName:  input.GymName,
		TeamName: input.TeamName,
		Active:   input.Active,
	}

	err := c.GymRepository.Create(&gym)
	if err != nil {
		return err
	}

	return nil

}

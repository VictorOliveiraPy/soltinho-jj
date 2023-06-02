package usecase


import "github.com/VictorOliveiraPy/internal/entity"

type CreateStudentUseCase struct {
	StudentRepository entity.StudentRepositoryInterface
}

func NewCreaStudentUseCase(
	StudentRepository entity.StudentRepositoryInterface,
) *CreateStudentUseCase {
	return &CreateStudentUseCase{
		StudentRepository: StudentRepository,
	}
}

type StudentInput struct {
	ID           string `json:"id"`
	GymID        string `json:"gym_id"`
	Name      	 string `json:"Name"`
	Graduation   string `json:"graduation"`
	Active       bool   `json:"active"`
	TrainingTime string `json:"training_time"`
}



func (c *CreateStudentUseCase) Execute(input StudentInput) error {
	student := entity.Student{
		ID:       input.ID,
		GymID: input.GymID,
		Name: input.Name,
		Graduation: input.Graduation,
		Active:   input.Active,
		TrainingTime: input.TrainingTime,
	}
	
	err := c.StudentRepository.Create(&student)
	if err != nil {
		return err
	}

	return nil

}

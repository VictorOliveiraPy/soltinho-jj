package db

import (
	"database/sql"

	"github.com/VictorOliveiraPy/internal/entity"
)


type StudentRepository struct {
	Db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{Db: db}
}

func (repo *StudentRepository) Create(student *entity.Student) error {
	stmt, err := repo.Db.Prepare("INSERT INTO students (id, gym_id, name, graduation, active, training_time) VALUES ($1, $2, $3, $4, $5, $6)")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(student.ID, student.GymID, student.Name, student.Graduation, student.Active, student.TrainingTime)
	if err != nil {
		return err
	}
	return nil
}
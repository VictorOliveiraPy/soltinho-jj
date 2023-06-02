package db

import (
	"database/sql"
	"github.com/VictorOliveiraPy/internal/entity"
)

type GymRepository struct {
	Db *sql.DB
}

func NewGymRepository(db *sql.DB) *GymRepository {
	return &GymRepository{Db: db}
}

func (repo *GymRepository) Create(gym *entity.Gym) error {
	stmt, err := repo.Db.Prepare("INSERT INTO gyms (id, user_id, gym_name, team_name, active) VALUES ($1, $2, $3, $4, $5)")

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(gym.ID, gym.UserID, gym.GymName, gym.TeamName, gym.Active)
	if err != nil {
		return err
	}
	return nil
}

func (repo *GymRepository) FindByName(gym_name string) (*entity.Gym, error) {
	query := "SELECT id, user_id, gym_name, team_name, active FROM gyms WHERE gym_name = $1 LIMIT 1"
	row := repo.Db.QueryRow(query, gym_name)
	var gym entity.Gym

	err := row.Scan(&gym.ID, &gym.UserID, &gym.GymName, &gym.TeamName, &gym.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &gym, nil
}

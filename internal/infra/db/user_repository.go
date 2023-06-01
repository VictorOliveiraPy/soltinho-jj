package db

import (
	"database/sql"

	"github.com/VictorOliveiraPy/internal/entity"
)


type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repo *UserRepository) Create(user *entity.User) error {
	stmt, err := repo.Db.Prepare("INSERT INTO users (id, username, password, email, role_id, active) VALUES ($1, $2, $3, $4, $5, $6)")

	println(stmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.UserName, user.Password, user.Email, user.RoleID, user.Active)
	if err != nil {
		return err
	}
	return nil
}

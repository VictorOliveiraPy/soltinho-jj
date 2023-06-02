package db

import (
	"database/sql"
	"fmt"

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

func (repo *UserRepository) FindByEmail(email string) (*entity.User, error) {
	query := "SELECT id, username, password, email, role_id, active FROM users WHERE email = $1"
	row := repo.Db.QueryRow(query, email)
	var user entity.User

	err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email, &user.RoleID, &user.Active)
	if err != nil {
		println("aqui")
		if err == sql.ErrNoRows {
			println("aqui2")
			return nil, nil
		}
		println("aqui 3")
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

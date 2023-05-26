// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, name, email, phone, academy_name, instructor_belt, password)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type CreateUserParams struct {
	ID             string
	Name           string
	Email          string
	Phone          string
	AcademyName    string
	InstructorBelt string
	Password       string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.AcademyName,
		arg.InstructorBelt,
		arg.Password,
	)
	return err
}

const findByEmail = `-- name: FindByEmail :one
SELECT id, name, email, phone, academy_name, instructor_belt, password
FROM users
WHERE email = $1
`

func (q *Queries) FindByEmail(ctx context.Context, email string) (*User, error) {
	row := q.db.QueryRowContext(ctx, findByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.AcademyName,
		&i.InstructorBelt,
		&i.Password,
	)
	return &i, err
}

const getusers = `-- name: Getusers :one
SELECT id, name, email, phone, academy_name, instructor_belt, password FROM users WHERE id = $1
`

func (q *Queries) Getusers(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getusers, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.AcademyName,
		&i.InstructorBelt,
		&i.Password,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, email, phone, academy_name, instructor_belt, password FROM users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Phone,
			&i.AcademyName,
			&i.InstructorBelt,
			&i.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET name = $2,
    email = $3,
    phone = $4,
    academy_name = $5,
    instructor_belt = $6,
    password = $7
WHERE id = $1
`

type UpdateUserParams struct {
	ID             string
	Name           string
	Email          string
	Phone          string
	AcademyName    string
	InstructorBelt string
	Password       string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.AcademyName,
		arg.InstructorBelt,
		arg.Password,
	)
	return err
}


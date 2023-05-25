-- name: Getusers :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users;


-- name: CreateUser :exec
INSERT INTO users (id, name, email, phone, academy_name, instructor_belt, password)
	VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateUser :exec
UPDATE users
SET name = $2,
    email = $3,
    phone = $4,
    academy_name = $5,
    instructor_belt = $6,
    password = $7
WHERE id = $1;


-- name: FindByEmail :one
SELECT *
FROM users
WHERE email = $1;

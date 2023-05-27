-- name: CreateStudent :exec
INSERT INTO students (id, gym_id, name, graduation, active, training_time)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: CreateGym :exec
INSERT INTO gyms (id, user_id, gym_name, team_name, active)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetGymByID :one
SELECT id, user_id, gym_name, team_name, active
FROM gyms
WHERE id = $1
LIMIT 1;


-- name: GetAllGyms :many
SELECT * FROM gyms;



-- name: GetByGymName :one
SELECT * FROM gyms WHERE gym_name = $1 LIMIT 1;

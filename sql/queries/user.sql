-- name: CreateUser :exec
INSERT INTO users (id, username, password, email, role_id, active)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET username = $2,
    password = $3,
    email = $4,
    role_id = $5,
    active = $6
WHERE id = $1;

-- name: GetUserRole :one
SELECT u.id, u.username, r.name AS role
FROM users u
JOIN user_roles r ON u.role_id = r.id
WHERE u.id = $1;


-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;


-- name: GetUserRoleName :one
SELECT ur.name FROM user_roles ur JOIN users u ON u.role_id = ur.id WHERE u.id = $1;


-- name: GetUserCompleteProfile :one
SELECT u.id, u.username, u.email, u.role_id, u.active,
       g.id AS gym_id, g.gym_name, g.team_name, g.active AS gym_active,
       s.id AS student_id, s.graduation, s.active AS student_active, s.training_time
FROM users u
LEFT JOIN gyms g ON u.id = g.user_id
LEFT JOIN students s ON g.id = s.gym_id
WHERE u.id = $1;

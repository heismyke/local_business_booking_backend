-- name: CreateUser :one
INSERT INTO users (
  name, email, phone, role,created_at,updated_at
) VALUES (
  $1, $2, $3, $4, NOW(), NOW()
) RETURNING *;


-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :exec
UPDATE users
  set name = $2,
  email = $3,
  phone = $4,
  role = $5
WHERE id = $1;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;
-- name: CreateBusinessHours :one
INSERT INTO business_hours (
  business_id, day_of_week, open_time, close_time
) VALUES (
  $1, $2, $3,$4
)RETURNING *;

-- name: GetBusinessHour :one
SELECT * FROM business_hours
WHERE id = $1 LIMIT 1;

-- name: ListBusinessHours :many
SELECT * FROM business_hours
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateBusinessHour :exec
UPDATE business_hours
  set business_id = $2,
  day_of_week = $3,
  open_time = $4,
  close_time = $5
WHERE id = $1;

-- name: DeleteBusinessHour :one
DELETE FROM business_hours
WHERE id = $1
RETURNING *;
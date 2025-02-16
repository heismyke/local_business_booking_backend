-- name: CreateBookings :one
INSERT INTO bookings (
  user_id, business_id, service, date, status
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetBooking :one
SELECT * FROM bookings
WHERE id = $1
LIMIT 1;

-- name: ListBookings :many
SELECT * FROM bookings
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateBooking :exec
UPDATE bookings
SET user_id = $2, 
    business_id = $3,
    service = $4,
    date = $5,
    status = $6
WHERE id = $1;

-- name: DeleteBooking :one
DELETE FROM bookings
WHERE id = $1
RETURNING *;

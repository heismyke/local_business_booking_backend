-- name: CreateReview :one
INSERT INTO reviews (
  user_id, business_id, rating, comment
) VALUES (
  $1, $2, $3, $4
)RETURNING *;

-- name: GetReview :one
SELECT * FROM reviews
WHERE id = $1 LIMIT 1;

-- name: ListReviews :many
SELECT * FROM reviews
ORDER BY Id 
LIMIT $1
OFFSET $2;

-- name: UpdateReview :exec
UPDATE reviews
  set user_id = $2, 
  business_id = $3,
  rating = $4,
  comment = $5
WHERE id = $1;

-- name: DeleteReview :one
DELETE FROM reviews
WHERE id = $1
RETURNING *;

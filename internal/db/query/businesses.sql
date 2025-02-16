-- name: CreateBusinesses :one
INSERT INTO businesses(
  owner,
  name,
 address,
 lattitude,
 longitude,
 phone,
 email,
 category,
 services
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)RETURNING *;

-- name: GetBusinesses :one
SELECT * FROM businesses
WHERE id = $1 LIMIT 1;

-- name: ListBusinesses :many
SELECT * FROM businesses
ORDER BY id
LIMIT  $1
OFFSET  $2;

-- name: UpdateBusiness :exec
UPDATE businesses
  set  owner = $2,
  name = $3,
  address = $4,
  lattitude = $5,
  longitude = $6,
  phone = $7,
  email = $8,
  category = $9,
  services = $10
WHERE id = $1;

-- name: DeleteBusiness :one
DELETE FROM businesses
WHERE id = $1
RETURNING *;
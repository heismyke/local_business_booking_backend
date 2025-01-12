// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: businesses.sql

package db

import (
	"context"
	"encoding/json"
)

const createBusinesses = `-- name: CreateBusinesses :one
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
)RETURNING id, owner, name, address, lattitude, longitude, phone, email, category, services, created_at, updated_at
`

type CreateBusinessesParams struct {
	Owner     int64           `json:"owner"`
	Name      string          `json:"name"`
	Address   string          `json:"address"`
	Lattitude float64         `json:"lattitude"`
	Longitude float64         `json:"longitude"`
	Phone     string          `json:"phone"`
	Email     string          `json:"email"`
	Category  string          `json:"category"`
	Services  json.RawMessage `json:"services"`
}

func (q *Queries) CreateBusinesses(ctx context.Context, arg CreateBusinessesParams) (Business, error) {
	row := q.db.QueryRowContext(ctx, createBusinesses,
		arg.Owner,
		arg.Name,
		arg.Address,
		arg.Lattitude,
		arg.Longitude,
		arg.Phone,
		arg.Email,
		arg.Category,
		arg.Services,
	)
	var i Business
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Name,
		&i.Address,
		&i.Lattitude,
		&i.Longitude,
		&i.Phone,
		&i.Email,
		&i.Category,
		&i.Services,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteBusiness = `-- name: DeleteBusiness :one
DELETE FROM businesses
WHERE id = $1
RETURNING id, owner, name, address, lattitude, longitude, phone, email, category, services, created_at, updated_at
`

func (q *Queries) DeleteBusiness(ctx context.Context, id int64) (Business, error) {
	row := q.db.QueryRowContext(ctx, deleteBusiness, id)
	var i Business
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Name,
		&i.Address,
		&i.Lattitude,
		&i.Longitude,
		&i.Phone,
		&i.Email,
		&i.Category,
		&i.Services,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBusinesses = `-- name: GetBusinesses :one
SELECT id, owner, name, address, lattitude, longitude, phone, email, category, services, created_at, updated_at FROM businesses
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBusinesses(ctx context.Context, id int64) (Business, error) {
	row := q.db.QueryRowContext(ctx, getBusinesses, id)
	var i Business
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Name,
		&i.Address,
		&i.Lattitude,
		&i.Longitude,
		&i.Phone,
		&i.Email,
		&i.Category,
		&i.Services,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listBusinesses = `-- name: ListBusinesses :many
SELECT id, owner, name, address, lattitude, longitude, phone, email, category, services, created_at, updated_at FROM businesses
ORDER BY id
LIMIT  $1
OFFSET  $2
`

type ListBusinessesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBusinesses(ctx context.Context, arg ListBusinessesParams) ([]Business, error) {
	rows, err := q.db.QueryContext(ctx, listBusinesses, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Business
	for rows.Next() {
		var i Business
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Name,
			&i.Address,
			&i.Lattitude,
			&i.Longitude,
			&i.Phone,
			&i.Email,
			&i.Category,
			&i.Services,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateBusiness = `-- name: UpdateBusiness :exec
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
WHERE id = $1
`

type UpdateBusinessParams struct {
	ID        int64           `json:"id"`
	Owner     int64           `json:"owner"`
	Name      string          `json:"name"`
	Address   string          `json:"address"`
	Lattitude float64         `json:"lattitude"`
	Longitude float64         `json:"longitude"`
	Phone     string          `json:"phone"`
	Email     string          `json:"email"`
	Category  string          `json:"category"`
	Services  json.RawMessage `json:"services"`
}

func (q *Queries) UpdateBusiness(ctx context.Context, arg UpdateBusinessParams) error {
	_, err := q.db.ExecContext(ctx, updateBusiness,
		arg.ID,
		arg.Owner,
		arg.Name,
		arg.Address,
		arg.Lattitude,
		arg.Longitude,
		arg.Phone,
		arg.Email,
		arg.Category,
		arg.Services,
	)
	return err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: reviews.sql

package db

import (
	"context"
)

const createReview = `-- name: CreateReview :one
INSERT INTO reviews (
  user_id, business_id, rating, comment
) VALUES (
  $1, $2, $3, $4
)RETURNING id, user_id, business_id, rating, comment, created_at
`

type CreateReviewParams struct {
	UserID     int64  `json:"user_id"`
	BusinessID int64  `json:"business_id"`
	Rating     int32  `json:"rating"`
	Comment    string `json:"comment"`
}

func (q *Queries) CreateReview(ctx context.Context, arg CreateReviewParams) (Review, error) {
	row := q.db.QueryRowContext(ctx, createReview,
		arg.UserID,
		arg.BusinessID,
		arg.Rating,
		arg.Comment,
	)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.BusinessID,
		&i.Rating,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}

const deleteReview = `-- name: DeleteReview :one
DELETE FROM reviews
WHERE id = $1
RETURNING id, user_id, business_id, rating, comment, created_at
`

func (q *Queries) DeleteReview(ctx context.Context, id int64) (Review, error) {
	row := q.db.QueryRowContext(ctx, deleteReview, id)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.BusinessID,
		&i.Rating,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}

const getReview = `-- name: GetReview :one
SELECT id, user_id, business_id, rating, comment, created_at FROM reviews
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetReview(ctx context.Context, id int64) (Review, error) {
	row := q.db.QueryRowContext(ctx, getReview, id)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.BusinessID,
		&i.Rating,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}

const listReviews = `-- name: ListReviews :many
SELECT id, user_id, business_id, rating, comment, created_at FROM reviews
ORDER BY Id 
LIMIT $1
OFFSET $2
`

type ListReviewsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListReviews(ctx context.Context, arg ListReviewsParams) ([]Review, error) {
	rows, err := q.db.QueryContext(ctx, listReviews, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Review
	for rows.Next() {
		var i Review
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.BusinessID,
			&i.Rating,
			&i.Comment,
			&i.CreatedAt,
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

const updateReview = `-- name: UpdateReview :exec
UPDATE reviews
  set user_id = $2, 
  business_id = $3,
  rating = $4,
  comment = $5
WHERE id = $1
`

type UpdateReviewParams struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	BusinessID int64  `json:"business_id"`
	Rating     int32  `json:"rating"`
	Comment    string `json:"comment"`
}

func (q *Queries) UpdateReview(ctx context.Context, arg UpdateReviewParams) error {
	_, err := q.db.ExecContext(ctx, updateReview,
		arg.ID,
		arg.UserID,
		arg.BusinessID,
		arg.Rating,
		arg.Comment,
	)
	return err
}

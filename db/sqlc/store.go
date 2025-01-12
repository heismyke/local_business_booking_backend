package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct{
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,	
		Queries: New(db),
	}
}

func (store *Store ) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil{
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

//TransferTxParams contains the input parameters of the transfer transaction
type CreateUserTxParams struct{
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  string `json:"role"`	
}

type CreateUserTxResult struct {
	UserID int64 `json:"user_id"`
}


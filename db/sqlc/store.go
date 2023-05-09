package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// Begin a transaction
	tx, err := store.db.BeginTx(ctx, nil)
	// incase of error, return error
	if err != nil {
		return err
	}
	// start executing the required queries
	q := New(tx)
	err = fn(q)
	//  incase of any error, rollback the transaction
	if err != nil {
		//  if there is a error with rolling back the transaction, error log this rollback
		if rbErr := tx.Rollback(); rbErr != nil {
			// log the error and the rollback error
			return fmt.Errorf("tx err: %v, rb.err: %v", err, rbErr)
		}
		// if there is no rollback error then just return error after rolling back
		return err
	}
	//  if there is no error while executing the functions, commit the transaction.
	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxparams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to another
// It creates a transfer record, add account entries, and update account's balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxparams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx,
		func(q *Queries) error {
			var err error

			result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
				FromAccountID: arg.FromAccountID,
				ToAccountID:   arg.ToAccountID,
				Amount:        arg.Amount,
			})
			if err != nil {
				return err
			}

			result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
				AccountID: arg.FromAccountID,
				Amount:    -arg.Amount,
			})
			if err != nil {
				return err
			}

			result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
				AccountID: arg.ToAccountID,
				Amount:    arg.Amount,
			})
			if err != nil {
				return err
			}
			// TODO: "update accounts" Balance
			return nil
		})

	return result, err
}

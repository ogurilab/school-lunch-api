package db

import (
	"context"
	"fmt"
)

func (query *SQLQuery) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := query.db.Begin()

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

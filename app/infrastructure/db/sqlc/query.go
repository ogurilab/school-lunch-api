package db

import (
	"context"
	"database/sql"

	"github.com/ogurilab/school-lunch-api/domain"
)

type Query interface {
	Querier
	CreateDishTx(ctx context.Context, dish *domain.Dish, menuID string) error
	CreateDishesTx(ctx context.Context, dishes []*domain.Dish, menuID string) error
}

type SQLQuery struct {
	*Queries
	db *sql.DB
}

func NewQuery(db *sql.DB) Query {
	return &SQLQuery{
		Queries: New(db),
		db:      db,
	}
}

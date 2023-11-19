package db

import "database/sql"

type Query interface {
	Querier
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

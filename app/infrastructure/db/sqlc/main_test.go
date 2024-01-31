package db

import (
	"database/sql"
	"os"
	"testing"

	"github.com/ogurilab/school-lunch-api/bootstrap"
)

type TestQuery interface {
	Query
	truncateMenusTable() error
	truncateCitiesTable() error
}

type query struct {
	Query
	db *sql.DB
}

var testQuery TestQuery

const (
	TEST_DB_URL = "root:root@tcp(localhost:3306)/school_lunch_test?charset=utf8mb4&parseTime=True"
)

func newTestQuery(db *sql.DB) TestQuery {
	return &query{
		Query: NewQuery(db),
		db:    db,
	}
}

func (q *query) truncateMenusTable() error {
	_, err := q.db.Exec("TRUNCATE TABLE menus")
	return err
}

func (q *query) truncateCitiesTable() error {
	_, err := q.db.Exec("TRUNCATE TABLE cities")

	return err
}

func TestMain(m *testing.M) {
	os.Setenv("DB_SOURCE", TEST_DB_URL)
	app := bootstrap.NewApp("../../../")

	testQuery = newTestQuery(app.DB)

	os.Exit(m.Run())
}

package db

import (
	"os"
	"testing"

	"github.com/ogurilab/school-lunch-api/bootstrap"
)

var testQuery Query

const (
	TEST_DB_URL = "root:root@tcp(localhost:3306)/school_lunch_test?charset=utf8mb4&parseTime=True"
)

func TestMain(m *testing.M) {
	os.Setenv("DB_SOURCE", TEST_DB_URL)
	app := bootstrap.NewApp("../../../")

	testQuery = NewQuery(app.DB)

	os.Exit(m.Run())
}

package db

import (
	"os"
	"testing"

	"github.com/ogurilab/school-lunch-api/bootstrap"
)

var testQuery Query

func TestMain(m *testing.M) {

	app := bootstrap.NewApp("../../../")

	testQuery = NewQuery(app.DB)

	os.Exit(m.Run())
}

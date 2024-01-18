package main

import (
	"github.com/ogurilab/school-lunch-api/bootstrap"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	"github.com/ogurilab/school-lunch-api/server"
)

func main() {
	app := bootstrap.NewApp(".")
	env := app.Env

	query := db.NewQuery(app.DB)
	bootstrap.RunMigration(env.MigrationURL, env.DBSource)
	defer bootstrap.CloseDatabase(app.DB)

	server.Run(env, query)

}

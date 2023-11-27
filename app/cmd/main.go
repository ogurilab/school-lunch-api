package main

import (
	"github.com/ogurilab/school-lunch-api/bootstrap"
	"github.com/ogurilab/school-lunch-api/server"
)

func main() {
	app := bootstrap.NewApp("../")
	env := app.Env

	server.Run(env)

}

package bootstrap

import (
	"database/sql"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Application struct {
	Env Env
	DB  *sql.DB
}

func NewApp(path string) (app Application) {

	env, err := NewEnv(path)

	if env.ENVIRONMENT == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load env")
	}

	db, err := NewDatabase(env)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	return Application{
		Env: env,
		DB:  db,
	}
}

package bootstrap

import (
	// mysql driver
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

const (
	maxLifTime = 5
	maxIdle    = 5
	maxOpen    = 5
)

type DB *sql.DB

func NewDatabase(env Env) (*sql.DB, error) {

	db, err := sql.Open("mysql", env.DBSource)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(maxLifTime)
	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	log.Info().Msg("Successfully connected to database")

	return db, nil
}

func CloseDatabase(db *sql.DB) {
	err := db.Close()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to disconnect from database")
	}

	log.Info().Msg("Successfully disconnected from database")

}

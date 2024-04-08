package model

import (
	"fmt"
	"sample/src/config"
	"sample/src/helper"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func Migrate() {
	// Migrate the schema
	err := godotenv.Load(".env")
	helper.CheckErr(err)
	settings := config.GetSettings()
	sqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", settings.DB_USERNAME, settings.DB_PASSWORD, settings.DB_SERVER, settings.DB_PORT, settings.DB_NAME)
	m, err := migrate.New("file://migrations/", sqlInfo)
	helper.CheckErr(err)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		helper.CheckErr(err)
	}
	log.Info().Msg("Migration completed")
}

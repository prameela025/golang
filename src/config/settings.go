package config

import (
    "os"
)

type Settings struct {
    DB_SERVER        string
    DB_PORT          string
    DB_NAME          string
    DB_USERNAME      string
    DB_PASSWORD      string
}

func GetSettings() *Settings {
    dbServer := os.Getenv("DATABASE_SERVER")
    dbPort := os.Getenv("DATABASE_PORT")
    dbName := os.Getenv("DATABASE_NAME")
    dbUsername := os.Getenv("DATABASE_USERNAME")
    dbPassword := os.Getenv("DATABASE_PASSWORD")

    return &Settings{
        DB_SERVER:        dbServer,
        DB_PORT:          dbPort,
        DB_NAME:          dbName,
        DB_USERNAME:      dbUsername,
        DB_PASSWORD:      dbPassword,
    }
}

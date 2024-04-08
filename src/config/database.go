package config

import (
	"fmt"
	"sample/src/helper"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func DatabaseConnection() *gorm.DB {
	err := godotenv.Load(".env")
	helper.CheckErr(err)
	settings := GetSettings()
	
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", settings.DB_SERVER, settings.DB_PORT, settings.DB_USERNAME, settings.DB_NAME, settings.DB_PASSWORD)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.CheckErr(err)
	return db
}

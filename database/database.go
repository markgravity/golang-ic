package database

import (
	"os"

	"github.com/markgravity/golang-ic/helpers/log"

	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDatabase(databaseURL string) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to %v database: %v", gin.Mode(), err)
	} else {
		viper.Set("database", db)
		log.Println(gin.Mode() + " database connected successfully.")
	}
	database = db

	migrateDB(db)
}

func migrateDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to convert gormDB to sqlDB: %v", err)
	}

	err = goose.Up(sqlDB, "database/migrations", goose.WithAllowMissing())
	if err != nil {
		log.Errorf("Failed to migrate database: %v", err)
	} else {
		log.Println("Migrated database successfully.")
	}
}

func GetDB() *gorm.DB {
	if database == nil {
		InitDatabase(GetDatabaseURL())
	}

	return database
}

func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}

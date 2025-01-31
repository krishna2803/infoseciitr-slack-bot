package database

import (
	"fmt"
	"os"

	"infosec/key-bot/models"
	"infosec/key-bot/pkg/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func connect() (*gorm.DB, error) {
	log.GetLogger().Info("connecting to database...")

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", db_host, db_user, db_pass, db_name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	log.GetLogger().Info("connected to database")

	err = db.AutoMigrate(&models.CTFEvent{}, &models.Key{}, &models.Member{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Init() error {
	db, err := connect()
	if err != nil {
		return err
	}
	DB = db
	return nil
}

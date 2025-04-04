package dal

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"whatwhen/models"
)

var Database *gorm.DB

func ConnectToPostgres() error {
	connectionStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		return err
	}
	err = database.AutoMigrate(&models.Task{})
	if err != nil {
		return err
	}

	Database = database
	return nil
}

package database

import (
	"fmt"
	"log"
	"os"
	"smartjob/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database  Connect Error : ", err)
	}
}

func AutoMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.RefreshToken{},
		&models.Announcement{},
		&models.JobPost{},
		&models.Question{},
		&models.QuestionOption{},
		&models.JobQuestion{},
		&models.Application{},
		&models.UserAnswer{},
	)
	if err != nil {
		log.Fatal("Auto Migrate Error", err)
	}
}

package database

import (
	"auth-service/models"
	"auth-service/utils"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		utils.GetEnv("DB_HOST", "localhost"),
		utils.GetEnv("DB_USER", "postgres"),
		utils.GetEnv("DB_PASSWORD", "postgres"),
		utils.GetEnv("DB_NAME", "skill-match-auth-db"),
		utils.GetEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	DB = db

	seedDummyUsers()
}

func seedDummyUsers() {
	var count int64
	DB.Model(&models.User{}).Count(&count)

	if count == 0 {
		log.Println("Seeding dummy users...")

		dummyUsers := []models.User{
			createUser("alice@example.com", "password123"),
			createUser("bob@example.com", "qwerty123"),
		}

		if err := DB.Create(&dummyUsers).Error; err != nil {
			log.Println("Failed to seed users:", err)
		} else {
			log.Println("Dummy users seeded.")
		}
	} else {
		log.Println("Users already exist, skipping seed.")
	}
}

func createUser(email, plainPassword string) models.User {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	return models.User{
		Email:    email,
		Password: string(hashed),
	}
}

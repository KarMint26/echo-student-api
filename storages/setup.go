package storages

import (
	"fmt"

	"github.com/KarMint26/echo-student-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define config struct
type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

// Open connection to database
func ConnectionDatabase(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// Migrate students table
	err = db.AutoMigrate(&models.Students{})
	if err != nil {
		return nil, err
	}

	DB = db

	return db, nil
}

func GetDB() *gorm.DB {
	return DB
}
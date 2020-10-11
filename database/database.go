package database

import (
	"github.com/fikrimohammad/Premier-League-DB/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Instance is a function to fetch database connection
func Instance() *gorm.DB {
	return db
}

// Connect is function to connect to a database
func Connect() {
	var err error
	dsn := "user=postgres password=postgres dbname=pl_development port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error when connecting to database")
	}
	db.AutoMigrate(&models.Position{})
	db.AutoMigrate(&models.Stadium{})
	db.AutoMigrate(&models.Club{})
}

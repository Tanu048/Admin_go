package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Tanu048/Admin_go/models"
)

// DB is the global database variable we will use in our routes
var DB *gorm.DB

func ConnectDb() {
	// STOP! Replace the string below with your actual Supabase URI
	dsn := "postgresql://postgres:[your-password]@[your-host]:6543/postgres"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("DB connection failed! Did you paste your Supabase URL? \n", err)
	}

	log.Println("Database connected successfully")

	// This automatically creates the Student and Mentor tables in Supabase
	err = db.AutoMigrate(&models.Admin{})
	if err != nil {
		log.Fatal("Migration failed: \n", err)
	}

	DB = db
}

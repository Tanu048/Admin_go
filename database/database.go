package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Tanu048/Admin_go/models"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := "postgresql://postgres:[your-password]@[your-host]:6543/postgres" //change later

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("DB connection failed! Did you paste your Supabase URL? \n", err)
	}

	log.Println("Database connected successfully")

	err = db.AutoMigrate(&models.Admin{})
	if err != nil {
		log.Fatal("Migration failed: \n", err)
	}

	DB = db
}

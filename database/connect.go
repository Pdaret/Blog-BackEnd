package database

import (
	"log"
	"os"

	"github.com/Sifouo/Blog-BackEnd/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Blog{})
	DB = database
}

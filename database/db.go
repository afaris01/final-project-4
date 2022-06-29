package database

import (
	"final-project-4/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var (
// 	host     = "localhost"
// 	user     = "root"
// 	password = ""
// 	dbPort   = "443"
// 	dbName   = "toko-belanja"
// 	err      error
// )

func MulaiDB() *gorm.DB{
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbName, dbPort)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal menyambung ke database :", err)
	}

	fmt.Println("Koneksi Sukses")
	database.Debug().AutoMigrate(models.User{}, models.Product{}, models.Category{})
	return database
}
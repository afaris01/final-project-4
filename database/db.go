package database

import (
	"final-project-4/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "root"
	password = ""
	dbPort   = "443"
	dbName   = "toko-belanja"
	err      error
)

func MulaiDB() *gorm.DB{
	dsn := "root@tcp(127.0.0.1:3306)/toko-belanja?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal menyambung ke database :", err)
	}

	fmt.Println("Koneksi Sukses")
	database.Debug().AutoMigrate(models.User{}, models.Product{}, models.Category{})
	return database
}
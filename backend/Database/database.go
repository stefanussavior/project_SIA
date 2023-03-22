package database

import (
	"backend/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Data .env tidak dapat di baca")
		return
	}

	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
	dsn, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("error", err.Error())
		return
	}
	dsn.AutoMigrate(models.Siswa{})
	fmt.Println("Database terkoneksi")
}

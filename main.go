package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		initializeDB()
	})
	return db
}

func initializeDB() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS") // or DB_PASS
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	fmt.Println("ENV:", dbUser, dbPassword, dbHost, dbPort, dbName) // debug

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	db = connection
}

func main() {

	initializeDB()

	db := GetDB()

	err := db.AutoMigrate(&entity.Admin{})

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Migration done!")
}

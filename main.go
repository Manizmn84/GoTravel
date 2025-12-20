package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
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
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	fmt.Println("ENV:", dbUser, dbPassword, dbHost, dbPort, dbName)

	// PostgreSQL DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to PostgreSQL: " + err.Error())
	}

	db = connection
}

func main() {
	initializeDB()

	db := GetDB()

	err := db.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.Airport{},
		&entity.FareClass{},
		&entity.Company{},
		&entity.Passenger{},
		&entity.Payment{},
		&entity.Reserve{},
		&entity.Route{},
		&entity.Seat{},
		&entity.Trip{},
	)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Migration done!")

	passengerRepo := postgress.NewPassengerRepository(db)
	passengerService := service.NewPassengerService(passengerRepo)

	reservationRepo := postgress.NewReservationRepository(db)
	reservationService := service.NewReservationService(reservationRepo, passengerRepo)

	paymentRepo := postgress.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo, reservationRepo)

	companyRepo := postgress.NewCompanyRepository(db)
	companyService := service.NewCompanyService(companyRepo)

	airportRepo := postgress.NewAirportRepository(db)
	airportService := service.NewAirportService(airportRepo)

	fareClassRepo := postgress.NewFareClassRepository(db)
	fareClassService := service.NewFareClassService(fareClassRepo)

	e := echo.New()

	e.POST("/CreatePassenger", passengerService.CreatePassenger)
	e.POST("/CreateReservation", reservationService.CreateReservation)
	e.POST("/CreatePayment", paymentService.CreatePayment)
	e.POST("/CreateCompany", companyService.CreateCompany)
	e.POST("/CreateAirport", airportService.CreateAirport)
	e.POST("/CreateFareClass", fareClassService.CreateFareClass)

	e.Start(":8080")
}

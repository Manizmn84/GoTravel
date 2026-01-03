package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
	"github.com/Manizmn84/GoTravel/internal/presentation/controller"
	"github.com/Manizmn84/GoTravel/internal/presentation/routes/http"
	"github.com/joho/godotenv"
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

	// Repositories
	passengerRepo := postgress.NewPassengerRepository(db)
	reservationRepo := postgress.NewReservationRepository(db)
	paymentRepo := postgress.NewPaymentRepository(db)
	companyRepo := postgress.NewCompanyRepository(db)
	airportRepo := postgress.NewAirportRepository(db)
	fareClassRepo := postgress.NewFareClassRepository(db)

	// Services
	passengerService := service.NewPassengerService(passengerRepo)
	reservationService := service.NewReservationService(reservationRepo, passengerRepo)
	paymentService := service.NewPaymentService(paymentRepo, reservationRepo)
	companyService := service.NewCompanyService(companyRepo)
	airportService := service.NewAirportService(airportRepo)
	fareClassService := service.NewFareClassService(fareClassRepo)

	// Controllers
	passengerController := controller.NewPassengerController(passengerService)
	reservationController := controller.NewReservationController(reservationService)
	paymentController := controller.NewPaymentController(paymentService)
	companyController := controller.NewCompanyController(companyService)
	airportController := controller.NewAirportController(airportService)
	fareClassController := controller.NewFareClassController(fareClassService)

	// Router
	r := http.NewRouter(
		passengerController,
		reservationController,
		paymentController,
		companyController,
		airportController,
		fareClassController,
	)

	// Start server
	r.Start(":8080")
}

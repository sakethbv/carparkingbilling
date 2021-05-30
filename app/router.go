package app

import (
	"carparkingbilling/app/kafka"
	"carparkingbilling/migrate"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// App Model
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

//InitializeDB - To establish connection to database
func (app *App) InitializeDB(dbname, dbhost, dbport, dbuser, dbpass, brokers, group, topic string) {

	dbURI := fmt.Sprintf("dbname=%s host=%s port=%s user=%s  password=%s  sslmode=disable",
		dbname, dbhost, dbport, dbuser, dbpass)

	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		log.Fatal("Could not connect to database", zap.Any("Error", err))
	}

	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(2)
	db.DB().SetConnMaxLifetime(5 * time.Minute)

	app.DB = migrate.DBMigrate(db)
	//Reading parking information from kafka and inserting into database
	go kafka.CreateKafkaConsumer(db, brokers, group, topic)
}

//Handlers for URLs
func (app *App) Handlers() *mux.Router {

	app.Router = mux.NewRouter().StrictSlash(true)

	//API to get the amount of time for which a car is parked
	app.Router.HandleFunc("/s1/parkingtime", app.GetParkingTime).Methods("GET")
	//API to get the amount to be paid for a parking slot based on duration
	app.Router.HandleFunc("/s1/parkingamount", app.GetParkingAmount).Methods("GET")

	return app.Router
}

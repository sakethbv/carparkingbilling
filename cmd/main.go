package main

import (
	router "carparkingbilling/app"
	"flag"
	"log"
	"net/http"

	"github.com/rs/cors"
)

//DBConfig - DB configuration
type DBConfig struct {
	DbName string `json:"db.name"`
	DbHost string `json:"db.host"`
	DbPort string `json:"db.port"`
	DbUser string `json:"db.user"`
	DbPass string `json:"db.pass"`
	Port   string `json:"port"`
}

var dbName = flag.String("db.name", "", "name of the database")
var dbHost = flag.String("db.host", "", "host where db is located")
var dbPort = flag.String("db.port", "", "port on which database is listening")
var dbUser = flag.String("db.user", "", "db user")
var dbPass = flag.String("db.pass", "", "password for database")
var port = flag.String("port", "", "listening port")
var kafkaBrokerUrl = flag.String("kafka.brokers", "localhost", "Kafka brokers seperated by comma")
var kafkaTopic = flag.String("kafka.topic", "billing", "Name of the Kafka topic")
var kafkaConsumerGroup = flag.String("kafka.consumergroup", "", "Name of Kafka consumer group")

func main() {

	flag.Parse()
	config := DBConfig{
		DbName: *dbName,
		DbHost: *dbHost,
		DbPort: *dbPort,
		DbUser: *dbUser,
		DbPass: *dbPass,
		Port:   *port,
	}

	app := router.App{}

	app.InitializeDB(config.DbName, config.DbHost, config.DbPort, config.DbUser, config.DbPass, *kafkaBrokerUrl, *kafkaConsumerGroup, *kafkaTopic)
	config.Port = ":" + config.Port
	http.Handle("/", app.Handlers())

	handler := cors.Default().Handler(app.Router)

	log.Fatal(http.ListenAndServe(config.Port, handler))
}

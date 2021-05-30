package kafka

import (
	"carparkingbilling/app/dboperation"
	"carparkingbilling/migrate"
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func CreateKafkaConsumer(DB *gorm.DB, Broker, Group, Topic string) error {
	//Creating kafka consumer
	//Kafkabrokers := strings.Split(Brokers, ",")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": Broker,
		"group.id":          Group,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return err
	}

	c.SubscribeTopics([]string{Topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			parkingMsg := migrate.ParkingData{}
			json.Unmarshal(msg.Value, &parkingMsg)

			err = dboperation.InsertKafkaMsgDB(DB, parkingMsg)
			if err != nil {
				fmt.Printf("Database insertion error: %v", err)
			}
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			return err
		}
	}
}

package dboperation

import (
	"carparkingbilling/migrate"
	"time"

	"github.com/jinzhu/gorm"
)

//GetParkingTimeDB func
func GetParkingTimeDB(DB *gorm.DB, ParkingID string) (ParkingData migrate.ParkingData, Err error) {

	if Err = DB.Where("parkingid = ?", ParkingID).Find(&ParkingData).Error; Err != nil {
		return ParkingData, Err
	}
	return ParkingData, nil
}

//InsertKafkaMsgDB func
func InsertKafkaMsgDB(DB *gorm.DB, KafkaParkingData migrate.ParkingData) (Err error) {

	time := time.Now().UTC()

	KafkaParkingData.TsCreated = time
	KafkaParkingData.ParkingTime = int(KafkaParkingData.OutTime.Sub(KafkaParkingData.InTime).Hours())

	//Inserting Kafka message into database
	if Err = DB.FirstOrCreate(&KafkaParkingData).Error; Err != nil {
		return Err
	}
	return nil
}

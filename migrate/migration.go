package migrate

import (
	"time"

	"github.com/jinzhu/gorm"
)

//ParkingData model
type ParkingData struct {
	ParkingID   string    `gorm:"column:parkingid" json:"parkingID,omitempty"`
	CarNumber   string    `gorm:"column:carnumber" json:"carNumber,omitempty"`
	InTime      time.Time `gorm:"column:intime" json:"inTime,omitempty"`
	OutTime     time.Time `gorm:"column:outtime" json:"outTime,omitempty"`
	ParkingTime int       `gorm:"column:parkingtime" json:"parkingTime,omitempty"` //In Hours
	TsCreated   time.Time `gorm:"column:tscreated" json:"tsCreated,omitempty"`
}

//TableName function for ParkingData
func (pd *ParkingData) TableName() string {
	return "parkingdata"
}

func DBMigrate(db *gorm.DB) *gorm.DB {

	db.AutoMigrate(&ParkingData{})

	return db
}

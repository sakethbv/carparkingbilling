package app

import (
	"carparkingbilling/app/dboperation"
	"carparkingbilling/app/util"
	"errors"
	"net/http"
)

//constant for parking charge per hour
const (
	AmountPerHour = 100
)

//GetParkingTime controller
func (app *App) GetParkingTime(ResponseWriter http.ResponseWriter, Request *http.Request) {

	parkingID := Request.FormValue("parkingId")

	var err error
	if len(parkingID) == 0 {
		util.RespondJSON(ResponseWriter, http.StatusBadRequest, errors.New("parkingId is mandatory").Error())
		return
	}
	//Retrieving the parking data based on the parkingID
	parkingData, err := dboperation.GetParkingTimeDB(app.DB, parkingID)
	if err != nil {
		util.RespondJSON(ResponseWriter, http.StatusOK, err.Error())
		return
	}
	//Sending the required fields in the JSON response
	var response = map[string]interface{}{
		"parkingId":  parkingData.ParkingID,
		"parkedTime": parkingData.ParkingTime,
	}
	util.RespondJSON(ResponseWriter, http.StatusOK, response)
}

//GetParkingAmount controller
func (app *App) GetParkingAmount(ResponseWriter http.ResponseWriter, Request *http.Request) {

	parkingID := Request.FormValue("parkingId")

	var err error
	if len(parkingID) == 0 {
		util.RespondJSON(ResponseWriter, http.StatusBadRequest, errors.New("parkingId is mandatory").Error())
		return
	}

	//Reusing the GetParkingTimeDB function to get the parking data
	parkingData, err := dboperation.GetParkingTimeDB(app.DB, parkingID)
	if err != nil {
		util.RespondJSON(ResponseWriter, http.StatusNotFound, err.Error())
		return
	}
	//Calculating the parking amount
	parkingAmount := parkingData.ParkingTime * AmountPerHour

	var response = map[string]interface{}{
		"parkingId":     parkingID,
		"parkingAmount": parkingAmount,
	}
	util.RespondJSON(ResponseWriter, http.StatusOK, response)
}

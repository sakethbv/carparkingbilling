package util

import (
	"encoding/json"
	"net/http"
)

//ResponseData model for displaying the status
type ResponseData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

//RespondJSON return the http response in json format
func RespondJSON(W http.ResponseWriter, Status int, Payload interface{}) {
	var res ResponseData
	res.Status = Status
	res.Data = Payload

	response, err := json.Marshal(res)
	if err != nil {
		W.WriteHeader(http.StatusInternalServerError)
		W.Write([]byte(err.Error()))
		return
	}
	W.Header().Set("Content-Type", "application/json")
	W.WriteHeader(Status)
	W.Write(response)
}

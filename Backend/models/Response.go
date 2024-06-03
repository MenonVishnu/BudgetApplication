package models

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string, errors interface{}) {
	var response Response

	response.Message = message
	response.Error = errors
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func SuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	var response Response

	response.Message = message
	response.Data = data
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

/*   **** NOTE ****
In the response.data and response.error section we are using interface beacause we don't
know what type of data will it hold. thus we can hold either a map or slice or array inside
of this interface
*/

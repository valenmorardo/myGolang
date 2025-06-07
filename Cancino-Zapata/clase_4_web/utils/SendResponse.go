package utils

import (
	"encoding/json"
	"net/http"

	"clase_4_web/structures"
)

func SendResponse(res http.ResponseWriter, statusCode int, success bool, message string, data any, errDetail any) {
	response := structures.Response{
		Success: success,
		Message: message,
		Data:    data,
		Error:   errDetail,
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	json.NewEncoder(res).Encode(response)
}

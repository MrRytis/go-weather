package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func JSON(w http.ResponseWriter, statusCode int, body interface{}) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprint(len(jsonData)))
	w.WriteHeader(statusCode)

	// Write the JSON data to the response
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Error writing JSON to response body", http.StatusInternalServerError)
		return
	}
}

func ErrorJSON(w http.ResponseWriter, message string, statusCode int) {
	err := Error{
		Message: message,
		Code:    statusCode,
	}

	JSON(w, statusCode, err)
}

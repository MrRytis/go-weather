package httpUtils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
}

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
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

	// Write the JSON data to the response body
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Error writing JSON to response body", http.StatusInternalServerError)
		return
	}
}

func ErrorJSON(w http.ResponseWriter, message string, statusCode int, errors ...Error) {
	err := ErrorResponse{
		Message: message,
		Code:    statusCode,
		Errors:  errors,
	}

	JSON(w, statusCode, err)
}

func ParseJSON(r *http.Request, w http.ResponseWriter, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		ErrorJSON(w, "Failed to parse request body", http.StatusBadRequest)
		return err
	}

	return err
}

package httpUtils

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

func XML(w http.ResponseWriter, statusCode int, body interface{}) {
	xmlData, err := xml.MarshalIndent(body, "", "    ")
	if err != nil {
		http.Error(w, "Failed to generate XML", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/xml")
	w.Header().Set("Content-Length", fmt.Sprint(len(xmlData)))
	w.WriteHeader(statusCode)

	// Write the JSON data to the response body
	_, err = w.Write(xmlData)
	if err != nil {
		http.Error(w, "Error writing XML to response body", http.StatusInternalServerError)
		return
	}
}

func ParseXML(r *http.Request, w http.ResponseWriter, v interface{}) {
	decoder := xml.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		ErrorJSON(w, "Failed to parse request body", http.StatusBadRequest)
	}
}

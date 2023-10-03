package handler

import (
	"github.com/MrRytis/go-weather/pkg/httpUtils"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/template/index.html")
	if err != nil {
		log.Println(err)
		httpUtils.ErrorJSON(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to text/html
	w.Header().Set("Content-Type", "text/html")

	// Render the HTML template with the provided data and write it to the response
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		httpUtils.ErrorJSON(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

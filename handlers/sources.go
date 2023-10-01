package handlers

import (
	"articlesfeedapi/domain"
	"encoding/json"
	"net/http"
)

func GetSourcesHandler(w http.ResponseWriter, r *http.Request) {
	sources := []domain.Source{
		{ID: 0, Name: "Mundo Deportivo"},
		{ID: 1, Name: "Diario AS"},
		{ID: 2, Name: "Marca"},
		{ID: 3, Name: "Oficial"},
		{ID: 4, Name: "90min"},
		{ID: 5, Name: "Sport"},
		{ID: 6, Name: "Sky Sports"},
		{ID: 7, Name: "El Desmarque"},
	}

	// Marshal the articles slice into JSON
	sourcesJSON, err := json.Marshal(sources)
	if err != nil {
		http.Error(w, "Failed to marshal sources into JSON", http.StatusInternalServerError)
		return
	}

	// Set appropriate content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(sourcesJSON)
}

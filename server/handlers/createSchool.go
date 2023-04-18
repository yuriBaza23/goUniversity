package handlers

import (
	"encoding/json"
	"fmt"
	"gopi/models"
	"gopi/repositories"
	"net/http"
)

func CreateSchool(w http.ResponseWriter, r *http.Request) {
	var school models.School

	err := json.NewDecoder(r.Body).Decode(&school)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := repositories.InsertSchool(school)

	var resp map[string]any
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": "Error inserting school",
		}
	} else {
		resp = map[string]any{
			"error":   false,
			"message": fmt.Sprintf(`School created with ID: %d`, id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

package handlers

import (
	"encoding/json"
	"gopi/internal/api/v1/repositories"
	"net/http"
)

func GetAllSchool(w http.ResponseWriter, r *http.Request) {
	schools, err := repositories.GetAllSchools()

	if err != nil {
		http.Error(w, "Não foi possível obter as instituições de ensino", http.StatusBadRequest)
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(schools)
	}
}

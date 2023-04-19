package handlers

import (
	"encoding/json"
	"gopi/internal/api/v1/models"
	"gopi/internal/api/v1/repositories"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateSchool(w http.ResponseWriter, r *http.Request) {
	var input models.School

	vars := mux.Vars(r)
	id, errConv := strconv.Atoi(vars["id"])
	if errConv != nil {
		http.Error(w, "Não foi possível converter o id", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	school, err := repositories.GetSchool(int64(id))
	if err != nil {
		http.Error(w, "Não foi possível obter a instituição de ensino", http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		input.Name = school.Name
	}
	if input.Type == "" {
		input.Type = school.Type
	}
	if input.Logo == "" {
		input.Logo = school.Logo
	}

	_, err = repositories.UpdateSchool(int64(id), input)
	if err != nil {
		http.Error(w, "Não foi possível atualizar a instituição de ensino", http.StatusBadRequest)
		return
	}

	school, err = repositories.GetSchool(int64(id))
	if err != nil {
		http.Error(w, "Não foi possível obter a instituição de ensino", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(school)
}

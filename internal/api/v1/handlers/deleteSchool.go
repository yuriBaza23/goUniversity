package handlers

import (
	"encoding/json"
	"gopi/internal/api/v1/repositories"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteSchool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, errConv := strconv.Atoi(vars["id"])
	if errConv != nil {
		http.Error(w, "Não foi possível converter o id", http.StatusBadRequest)
		return
	}

	school, err := repositories.DeleteSchool(int64(id))

	if err != nil {
		http.Error(w, "Não foi possível deletar a instituição de ensino", http.StatusBadRequest)
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(school)
	}
}

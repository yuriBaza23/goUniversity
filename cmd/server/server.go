package server

import (
	"fmt"
	"gopi/internal/api/v1/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func HttpInit(port string) {
	r := mux.NewRouter()

	r.HandleFunc("/school", handlers.GetAllSchool).Methods("GET")
	r.HandleFunc("/school/{id}", handlers.GetSchool).Methods("GET")
	r.HandleFunc("/school", handlers.CreateSchool).Methods("POST")
	r.HandleFunc("/school/{id}", handlers.DeleteSchool).Methods("DELETE")
	r.HandleFunc("/school/{id}", handlers.UpdateSchool).Methods("PUT")

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

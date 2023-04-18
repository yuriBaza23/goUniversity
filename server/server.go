package server

import (
	"fmt"
	"gopi/server/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Init(port string) {
	r := chi.NewRouter()

	r.Post("/school", handlers.CreateSchool)

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

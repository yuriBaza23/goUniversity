package server

import (
	"fmt"
	"gopi/internal/api/v1/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HttpInit(port string) {
	r := chi.NewRouter()

	r.Post("/school", handlers.CreateSchool)

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

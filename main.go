package main

import (
	"fmt"
	"gopi/config"
	"gopi/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/school", handlers.Create)

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetAPIConfig()), r)
}

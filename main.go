package main

import (
	"ex2/handlers"
	// "ex2/repositories"
	"ex2/storages"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.Printf("Starting up on http://localhost:3333\n")
	r := chi.NewRouter()
	dbManager, err := storages.NewPSQLManager()
	if err != nil {
		log.Printf("Error happened in database server marshal. Err: %s", err)
		return
	}
	h, err1 := handlers.NewUser(dbManager)
	if err1 != nil {
		log.Printf("Error happened in database server marshal. Err: %s", err)
		return
	}
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Route("/api/me", func(r chi.Router) {
		r.Use(h.AuthenHeader)
		r.Get("/", h.ResInformation)
	})
	http.ListenAndServe(":3333", r)
}

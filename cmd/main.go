package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kinesocks/cat_facts_service/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/breeds", handlers.GetBreeds)
	r.Get("/fact", handlers.GetFact)
	r.Get("/facts", handlers.GetFacts)

	fmt.Println("Server is running...")

	err := http.ListenAndServe("0.0.0.0:8888", r)
	if err != nil {
		log.Fatal(err)
	}

}

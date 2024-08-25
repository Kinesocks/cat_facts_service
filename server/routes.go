package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Kinesocks/cat_facts_service/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/breeds", s.GetBreeds)
	r.Get("/fact", s.GetFact)
	r.Get("/facts", s.GetFacts)

	return r
}

func (s *Server) GetBreeds(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	var query map[string]interface{}
	err = json.Unmarshal(b, &query)
	if err != nil {
		log.Fatal(err)
	}

	_, ok := query["limit"]

	if !ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	lim := int(query["limit"].(float64))

	resp := api.BreedsQuery(lim, s.ApiToken)

	var apiresp map[string]interface{}
	err = json.Unmarshal(resp, &apiresp)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiresp["data"])
}

func (s *Server) GetFact(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	var query map[string]interface{}
	err = json.Unmarshal(b, &query)
	if err != nil {
		log.Fatal(err)
	}

	_, ok := query["max_length"]

	if !ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	max_length := int(query["max_length"].(float64))

	resp := api.GetFactQuery(max_length, s.ApiToken)

	var apiresp map[string]interface{}
	err = json.Unmarshal(resp, &apiresp)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiresp["fact"])
}

func (s *Server) GetFacts(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	var query map[string]interface{}
	err = json.Unmarshal(b, &query)
	if err != nil {
		log.Fatal(err)
	}

	_, ok1 := query["max_length"]
	_, ok2 := query["limit"]

	if !ok1 || !ok2 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	max_length := int(query["max_length"].(float64))
	lim := int(query["limit"].(float64))

	resp := api.GetFactsQuery(max_length, lim, s.ApiToken)

	var apiresp map[string]interface{}
	err = json.Unmarshal(resp, &apiresp)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiresp["data"])
}

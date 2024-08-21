package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Kinesocks/cat_facts_service/api"
)

func GetFacts(w http.ResponseWriter, r *http.Request) {
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

	resp := api.GetFactsQuery(max_length, lim)

	var apiresp map[string]interface{}
	err = json.Unmarshal(resp, &apiresp)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiresp["data"])
}
